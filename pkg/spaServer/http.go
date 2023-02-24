package server

import (
	"log"
	"net/http"

	"github.com/vault-thirteen/MIME"
	hdr "github.com/vault-thirteen/header"
)

const IndexPageHdrContentType = mime.TypeTextHtml + "; charset=UTF-8"

func (srv *Server) httpRouter(rw http.ResponseWriter, req *http.Request) {
	/* DEBUG */
	//log.Println(pretty.Sprint(req.URL))
	//log.Println(pretty.Sprint(req.Proto))
	//log.Println(pretty.Sprint(req.Header))

	// Path is '/'.
	if len(req.URL.Path) == 1 {
		srv.respondWithData(rw, srv.cachedIndexPage, IndexPageHdrContentType)
		return
	}

	// Cached files ?
	path := req.URL.Path[1:]
	for _, cachedPage := range srv.cachedPages {
		if path == cachedPage.Name {
			srv.respondWithData(rw, cachedPage.Contents, cachedPage.MimeType)
			return
		}
	}

	// Otherwise, it is the index page.
	srv.respondWithData(rw, srv.cachedIndexPage, IndexPageHdrContentType)
}

func (srv *Server) respondWithData(
	rw http.ResponseWriter,
	data []byte,
	mimeType string,
) {
	rw.Header().Set(hdr.HttpHeaderContentType, mimeType)
	rw.Header().Set(hdr.HttpHeaderServer, ServerName)

	// CORS support.
	if len(srv.settings.AllowedOriginForCORS) > 0 {
		rw.Header().Set(hdr.HttpHeaderAccessControlAllowOrigin, srv.settings.AllowedOriginForCORS)
	}

	// 1.
	// If a request doesn't have an Authorization header, or you are already
	// using s-maxage or must-revalidate in the response, then you don't need
	// to use public.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control
	// 2.
	// If there is a Cache-Control header with the max-age or s-maxage
	// directive in the response, the Expires header is ignored.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expires
	rw.Header().Set(hdr.HttpHeaderCacheControl, srv.httpHdrCacheControl)

	rw.WriteHeader(http.StatusOK)

	_, err := rw.Write(data)
	if err != nil {
		log.Println(err)
	}
}
