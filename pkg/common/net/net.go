package net

import (
	"github.com/vault-thirteen/auxie/cio"
)

const (
	HostPortFormat = "%s:%d"
)

const (
	HttpHeaderFormat_CacheControl = "max-age=%d, must-revalidate"
)

func FormatHostPort(host string, port uint16) string {
	return cio.FormatMessage(HostPortFormat, host, port)
}

func FormatHttpHeader_CacheControl(httpCacheControlMaxAge uint) string {
	return cio.FormatMessage(HttpHeaderFormat_CacheControl, httpCacheControlMaxAge)
}
