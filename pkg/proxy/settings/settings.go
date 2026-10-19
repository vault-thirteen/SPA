package settings

import (
	"os"
	"strings"

	"github.com/vault-thirteen/auxie/boolean"
	"github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/auxie/number"
	"github.com/vault-thirteen/auxie/reader"

	"github.com/vault-thirteen/SPA/pkg/common/helper"
)

const (
	Err_FileIsNotSet                     = "file is not set"
	Err_ServerHostIsNotSet               = "server host is not set"
	Err_ServerPortIsNotSet               = "server port is not set"
	Err_ServerModeIsNotSet               = "server mode is not set"
	Err_ServerMode                       = "server mode error"
	Err_CertFileIsNotSet                 = "certificate file is not set"
	Err_KeyFileIsNotSet                  = "key file is not set"
	Err_HttpCacheControlMaxAge           = "HTTP cache control max-age error"
	Err_ProxyTargetServerAddressIsNotSet = "proxy target server address is not set"
	Err_IparcDbFileIsNotSet              = "IPARC database file is not set"
	Err_ForbiddenCountryCodes            = "forbidden country codes error"
)

const (
	ServerModeHttp    = "HTTP"
	ServerModeIdHttp  = 1
	ServerModeHttps   = "HTTPS"
	ServerModeIdHttps = 2
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

	// Where to redirect the requests.
	ProxyTargetServerAddress string

	// Path to IPARC database ZIP file.
	IPARCDbFile string

	// Allow unknown countries.
	AllowUnknownCountries bool

	// Forbidden Country Codes.
	ForbiddenCountryCodes []string

	// Is this server main proxy ?
	// Main proxy server does not return CORS headers.
	IsMainProxy bool
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
			err = errors.Combine(err, derr)
		}
	}()

	rdr := reader.New(file)
	var buf = make([][]byte, 12)

	for i := range buf {
		buf[i], err = rdr.ReadLineEndingWithCRLF()
		if err != nil {
			return stn, err
		}
	}

	// Server Host & Port.
	stn.ServerHost = strings.TrimSpace(string(buf[0]))

	stn.ServerPort, err = number.ParseUint16(strings.TrimSpace(string(buf[1])))
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

	stn.HttpCacheControlMaxAge, err = number.ParseUint(strings.TrimSpace(string(buf[5])))
	if err != nil {
		return stn, err
	}

	stn.AllowedOriginForCORS = strings.TrimSpace(string(buf[6]))

	// Proxy Target.
	stn.ProxyTargetServerAddress = strings.TrimSpace(string(buf[7]))

	// Path to IPARC database ZIP file.
	stn.IPARCDbFile = strings.TrimSpace(string(buf[8]))

	stn.AllowUnknownCountries, err = boolean.FromString(strings.TrimSpace(string(buf[9])))
	if err != nil {
		return stn, err
	}

	// Forbidden Country Codes.
	stn.ForbiddenCountryCodes = helper.ToUpperCase(helper.ParseCSV(string(buf[10])))

	stn.IsMainProxy, err = boolean.FromString(strings.TrimSpace(string(buf[11])))
	if err != nil {
		return stn, err
	}

	return stn, nil
}

func (stn *Settings) Check() (err error) {
	if len(stn.File) == 0 {
		return errors.ErrorFromText(Err_FileIsNotSet)
	}

	if len(stn.ServerHost) == 0 {
		return errors.ErrorFromText(Err_ServerHostIsNotSet)
	}

	if stn.ServerPort == 0 {
		return errors.ErrorFromText(Err_ServerPortIsNotSet)
	}

	if len(stn.ServerModeStr) == 0 {
		return errors.ErrorFromText(Err_ServerModeIsNotSet)
	} else {
		if (stn.ServerModeStr != ServerModeHttp) &&
			(stn.ServerModeStr != ServerModeHttps) {
			return errors.ErrorFromText(Err_ServerMode)
		}
	}

	if stn.ServerModeId == 0 {
		return errors.ErrorFromText(Err_ServerModeIsNotSet)
	} else {
		if (stn.ServerModeId != ServerModeIdHttp) &&
			(stn.ServerModeId != ServerModeIdHttps) {
			return errors.ErrorFromText(Err_ServerMode)
		}
	}

	switch stn.ServerModeStr {
	case ServerModeHttp:
		// Keys are not required.
	case ServerModeHttps:
		if len(stn.CertFile) == 0 {
			return errors.ErrorFromText(Err_CertFileIsNotSet)
		}
		if len(stn.KeyFile) == 0 {
			return errors.ErrorFromText(Err_KeyFileIsNotSet)
		}
	default:
		return errors.ErrorFromText(Err_ServerMode)
	}

	if stn.HttpCacheControlMaxAge == 0 {
		return errors.ErrorFromText(Err_HttpCacheControlMaxAge)
	}

	// AllowedOriginForCORS is not checked as it may be empty.

	// Proxy Target.
	if len(stn.ProxyTargetServerAddress) == 0 {
		return errors.ErrorFromText(Err_ProxyTargetServerAddressIsNotSet)
	}

	// Path to IPARC database ZIP file.
	if len(stn.IPARCDbFile) == 0 {
		return errors.ErrorFromText(Err_IparcDbFileIsNotSet)
	}

	if stn.ForbiddenCountryCodes == nil {
		return errors.ErrorFromText(Err_ForbiddenCountryCodes)
	}

	return nil
}
