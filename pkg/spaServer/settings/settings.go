package settings

import (
	"errors"
	"os"
	"strings"

	"github.com/vault-thirteen/SPA/pkg/common/helper"
	af "github.com/vault-thirteen/auxie/file"
	"github.com/vault-thirteen/auxie/reader"
	"github.com/vault-thirteen/errorz"
)

const (
	ErrFileIsNotSet           = "file is not set"
	ErrServerHostIsNotSet     = "server host is not set"
	ErrServerPortIsNotSet     = "server port is not set"
	ErrServerModeIsNotSet     = "server mode is not set"
	ErrServerMode             = "server mode error"
	ErrCertFileIsNotSet       = "certificate file is not set"
	ErrKeyFileIsNotSet        = "key file is not set"
	ErrHttpCacheControlMaxAge = "HTTP cache control max-age error"
)

const (
	ServerModeHttp    = "HTTP"
	ServerModeIdHttp  = 1
	ServerModeHttps   = "HTTPS"
	ServerModeIdHttps = 2
)

const (
	IndexHtmlPageFileName = "index.html"
)

// Settings is Server's settings.
type Settings struct {
	// Path to the File with these Settings.
	File string

	// Server's Host Name.
	ServerHost string

	// Server's Listen Port.
	ServerPort uint16

	// ServerMode is an HTTP mode selector.
	// Possible values are: HTTP and HTTPS.
	ServerModeStr string
	ServerModeId  byte

	// Server's Certificate and Key.
	CertFile string
	KeyFile  string

	// HttpCacheControlMaxAge is time in seconds for which this server's
	// response is fresh (valid). After this period clients will be refreshing
	// the stale content by re-requesting it from the server.
	HttpCacheControlMaxAge uint

	// Allowed Origin for cross-origin requests (CORS).
	AllowedOriginForCORS string

	// Names of files to be cached as pages.
	CachedPageFileNames []string
}

func NewSettingsFromFile(filePath string) (stn *Settings, err error) {
	stn = &Settings{
		File: filePath,
	}

	var file *os.File
	file, err = os.Open(stn.File)
	if err != nil {
		return stn, err
	}
	defer func() {
		derr := file.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	rdr := reader.New(file)
	var buf = make([][]byte, 8)

	for i := range buf {
		buf[i], err = rdr.ReadLineEndingWithCRLF()
		if err != nil {
			return stn, err
		}
	}

	// Server Host & Port.
	stn.ServerHost = strings.TrimSpace(string(buf[0]))

	stn.ServerPort, err = helper.ParseUint16(strings.TrimSpace(string(buf[1])))
	if err != nil {
		return stn, err
	}

	// Server Work Mode.
	stn.ServerModeStr = strings.ToUpper(strings.TrimSpace(string(buf[2])))
	switch stn.ServerModeStr {
	case ServerModeHttp:
		stn.ServerModeId = ServerModeIdHttp
	case ServerModeHttps:
		stn.ServerModeId = ServerModeIdHttps
	}

	// Certificate and Key for optional TLS.
	stn.CertFile = strings.TrimSpace(string(buf[3]))
	stn.KeyFile = strings.TrimSpace(string(buf[4]))

	stn.HttpCacheControlMaxAge, err = helper.ParseUint(strings.TrimSpace(string(buf[5])))
	if err != nil {
		return stn, err
	}

	stn.AllowedOriginForCORS = strings.TrimSpace(string(buf[6]))

	// Cached Page FileNames.
	tmp := strings.TrimSpace(string(buf[7]))
	tmpParts := strings.Split(tmp, ",")
	stn.CachedPageFileNames = make([]string, 0, len(tmpParts))

	var fileName string
	for _, tmpPart := range tmpParts {
		fileName = strings.TrimSpace(tmpPart)
		_, err = af.FileExists(fileName)
		if err != nil {
			return nil, err
		}

		stn.CachedPageFileNames = append(stn.CachedPageFileNames, fileName)
	}

	return stn, nil
}

func (stn *Settings) Check() (err error) {
	if len(stn.File) == 0 {
		return errors.New(ErrFileIsNotSet)
	}

	if len(stn.ServerHost) == 0 {
		return errors.New(ErrServerHostIsNotSet)
	}

	if stn.ServerPort == 0 {
		return errors.New(ErrServerPortIsNotSet)
	}

	if len(stn.ServerModeStr) == 0 {
		return errors.New(ErrServerModeIsNotSet)
	} else {
		if (stn.ServerModeStr != ServerModeHttp) &&
			(stn.ServerModeStr != ServerModeHttps) {
			return errors.New(ErrServerMode)
		}
	}

	if stn.ServerModeId == 0 {
		return errors.New(ErrServerModeIsNotSet)
	} else {
		if (stn.ServerModeId != ServerModeIdHttp) &&
			(stn.ServerModeId != ServerModeIdHttps) {
			return errors.New(ErrServerMode)
		}
	}

	switch stn.ServerModeStr {
	case ServerModeHttp:
		// Keys are not required.
	case ServerModeHttps:
		if len(stn.CertFile) == 0 {
			return errors.New(ErrCertFileIsNotSet)
		}
		if len(stn.KeyFile) == 0 {
			return errors.New(ErrKeyFileIsNotSet)
		}
	default:
		return errors.New(ErrServerMode)
	}

	if stn.HttpCacheControlMaxAge == 0 {
		return errors.New(ErrHttpCacheControlMaxAge)
	}

	// AllowedOriginForCORS is not checked as it may be empty.

	return nil
}
