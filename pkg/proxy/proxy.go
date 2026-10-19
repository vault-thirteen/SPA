package server

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/vault-thirteen/auxie/cio"
	ae "github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/auxie/header"
)

func (srv *Server) newProxy(targetHost string) (*httputil.ReverseProxy, error) {
	targetUrl, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(targetUrl)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		srv.modifyRequest(req)
	}

	proxy.ModifyResponse = srv.modifyResponse()
	proxy.ErrorHandler = srv.errorHandler()

	return proxy, nil
}

func (srv *Server) modifyRequest(req *http.Request) {
	req.Header.Set(HttpHeader_XProxy, ProxyServerName)
}

// modifyResponse is called when the target is reachable.
// This behaviour is hard-coded in the built-in library.
func (srv *Server) modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		if srv.settings.IsMainProxy {
			resp.Header.Del(header.HttpHeaderAccessControlAllowOrigin)
		}

		if resp.StatusCode == http.StatusOK {
			return nil
		}

		switch resp.StatusCode {
		case http.StatusBadRequest:
			return NewProxyError(http.StatusBadRequest, "bad request")

		case http.StatusInternalServerError:
			return NewProxyError(http.StatusInternalServerError, "internal server error")

		default:
			return NewProxyError(resp.StatusCode, "unknown error")
		}
	}
}

// errorHandler is called in two cases:
//  1. If the target is unreachable;
//  2. If the target is reachable, but ModifyResponse returns an error.
//
// This behaviour is hard-coded in the built-in library.
func (srv *Server) errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		var proxyError *ProxyError
		if errors.As(err, &proxyError) {
			// Target responded with a bad status code.
			// If it is our error, log it.
			if proxyError.StatusCode >= http.StatusInternalServerError {
				cio.LogMessage(proxyError.StatusCode, proxyError.Error())
			}
			w.WriteHeader(proxyError.StatusCode)
			return
		}

		var netError *net.OpError
		if errors.As(err, &netError) {
			// Target is unreachable.
			ae.LogError(netError)
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		// Unsupported error type.
		ae.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
