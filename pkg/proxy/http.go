package server

import (
	"log"
	"net/http"
)

// Custom HTTP Headers.
const (
	HttpHeaderXProxy             = "X-Proxy"
	HttpHeaderXClientCountryCode = "X-ClientCountryCode"
)

func (srv *Server) httpRouter(rw http.ResponseWriter, req *http.Request) {
	clientIPARange, err := srv.getClientIPv4AddressRangeNR(rw, req)
	if err != nil {
		log.Println(err)
		return
	}

	ok, clientCountryCode := srv.isIPARAllowed(clientIPARange)
	if !ok {
		srv.breakConnection(rw)
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

func (srv *Server) breakConnection(rw http.ResponseWriter) {
	hj, ok := rw.(http.Hijacker)
	if !ok {
		srv.respondWithInternalServerError(rw)
		return
	}

	conn, _, err := hj.Hijack()
	if err != nil {
		srv.respondWithInternalServerError(rw)
		return
	}

	err = conn.Close()
	if err != nil {
		srv.respondWithInternalServerError(rw)
		return
	}
}
