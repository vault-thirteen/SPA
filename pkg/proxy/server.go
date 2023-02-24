package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"sync"
	"sync/atomic"
	"time"

	"github.com/vault-thirteen/IPARC/common/helper"
	"github.com/vault-thirteen/IPARC/iparc"
	set "github.com/vault-thirteen/SPA/pkg/proxy/settings"
)

const (
	ProxyServerName = "SPARP"
)

type Server struct {
	settings *set.Settings

	// HTTP(S) server.
	listenDsn  string
	httpServer *http.Server

	// Channel for an external controller. When a message comes from this
	// channel, a controller must stop this server. The server does not stop
	// itself.
	mustBeStopped chan bool

	// Internal control structures.
	subRoutines *sync.WaitGroup
	mustStop    *atomic.Bool
	httpErrors  chan error

	// HTTP header values.
	httpHdrCacheControl string

	// Proxy Server.
	proxy *httputil.ReverseProxy

	iparc *iparc.IPAddressV4RangeCollection
}

func NewServer(stn *set.Settings) (srv *Server, err error) {
	err = stn.Check()
	if err != nil {
		return nil, err
	}

	srv = &Server{
		settings:      stn,
		listenDsn:     fmt.Sprintf("%s:%d", stn.ServerHost, stn.ServerPort),
		mustBeStopped: make(chan bool, 2),
		subRoutines:   new(sync.WaitGroup),
		mustStop:      new(atomic.Bool),
		httpErrors:    make(chan error, 8),

		httpHdrCacheControl: fmt.Sprintf("max-age=%d, must-revalidate",
			stn.HttpCacheControlMaxAge),
	}
	srv.mustStop.Store(false)

	srv.proxy, err = srv.newProxy(srv.settings.ProxyTargetServerAddress)
	if err != nil {
		return nil, err
	}

	// HTTP Server.
	srv.httpServer = &http.Server{
		Addr:    srv.listenDsn,
		Handler: http.Handler(http.HandlerFunc(srv.httpRouter)),
	}

	// IPARC.
	err = srv.initIPARC()
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func (srv *Server) initIPARC() (err error) {
	var dbFilePath string
	dbFilePath, err = helper.UnpackDbFile(srv.settings.IPARCDbFile)
	if err != nil {
		return err
	}

	srv.iparc, err = iparc.NewFromCsvFile(dbFilePath)
	if err != nil {
		return err
	}

	return nil
}

func (srv *Server) GetListenDsn() (dsn string) {
	return srv.listenDsn
}

func (srv *Server) GetWorkMode() (modeId byte) {
	return srv.settings.ServerModeId
}

func (srv *Server) GetStopChannel() *chan bool {
	return &srv.mustBeStopped
}

func (srv *Server) Start() (err error) {
	srv.startHttpServer()

	srv.subRoutines.Add(1)
	go srv.listenForHttpErrors()

	return nil
}

func (srv *Server) Stop() (err error) {
	srv.mustStop.Store(true)

	ctx, cf := context.WithTimeout(context.Background(), time.Minute)
	defer cf()
	err = srv.httpServer.Shutdown(ctx)
	if err != nil {
		return err
	}

	close(srv.httpErrors)

	srv.subRoutines.Wait()

	return nil
}

func (srv *Server) startHttpServer() {
	go func() {
		var listenError error
		switch srv.settings.ServerModeId {
		case set.ServerModeIdHttp:
			listenError = srv.httpServer.ListenAndServe()
		case set.ServerModeIdHttps:
			listenError = srv.httpServer.ListenAndServeTLS(srv.settings.CertFile, srv.settings.KeyFile)
		}
		if (listenError != nil) && (listenError != http.ErrServerClosed) {
			srv.httpErrors <- listenError
		}
	}()
}

func (srv *Server) listenForHttpErrors() {
	defer srv.subRoutines.Done()

	for err := range srv.httpErrors {
		log.Println("Server error: " + err.Error())
		srv.mustBeStopped <- true
	}

	log.Println("HTTP error listener has stopped.")
}
