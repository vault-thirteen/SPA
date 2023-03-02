package server

import (
	"log"
	"net"
	"net/http"
	"time"
)

// Custom HTTP Headers.
const (
	HttpHeaderXProxy             = "X-Proxy"
	HttpHeaderXClientCountryCode = "X-ClientCountryCode"
)

const BCST = time.Millisecond * 50

func (srv *Server) httpRouter(rw http.ResponseWriter, req *http.Request) {
	clientIPARange, err := srv.getClientIPv4AddressRangeNR(rw, req)
	if err != nil {
		log.Println(err)
		return
	}

	ok, clientCountryCode := srv.isIPARAllowed(clientIPARange)
	if !ok {
		cerr := srv.breakConnection(rw)
		if cerr != nil {
			log.Println(cerr)
			srv.respondWithInternalServerError(rw)
		}
		return
	}

	req.Header.Set(HttpHeaderXClientCountryCode, clientCountryCode)

	srv.proxy.ServeHTTP(rw, req)
}

func (srv *Server) respondWithInternalServerError(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
}

func (srv *Server) respondWithNotImplemented(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNotImplemented)
}

func (srv *Server) breakConnection(rw http.ResponseWriter) (err error) {
	rc := http.NewResponseController(rw)

	var conn net.Conn
	conn, _, err = rc.Hijack()
	if err != nil {
		err = rc.SetWriteDeadline(time.Now().Add(time.Microsecond))
		if err != nil {
			return err
		}

		time.Sleep(BCST)

		return nil
	}

	err = conn.Close()
	if err != nil {
		return err
	}

	return nil
}
