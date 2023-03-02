package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/vault-thirteen/IPARC/ipa"
	"github.com/vault-thirteen/IPARC/ipad/country"
	"github.com/vault-thirteen/IPARC/ipar"
)

// getClientIPv4AddressRangeNR tries to find the IP v4 address range of the
// provided IP address. If an error occurs, it responds with an appropriate HTTP
// status code. The caller of this function does not need to respond.
func (srv *Server) getClientIPv4AddressRangeNR(
	rw http.ResponseWriter,
	req *http.Request,
) (ipaRange *ipar.IPAddressV4Range, err error) {
	// Check the IP address of the client.
	clientHost, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		srv.respondWithInternalServerError(rw)
		return nil, err
	}

	clientIPAddr := net.ParseIP(clientHost)
	if clientIPAddr == nil {
		srv.respondWithInternalServerError(rw)
		return nil, fmt.Errorf("client's host is not an IP address: %v", clientHost)
	}

	var clientIPA ipa.IPAddressV4
	clientIPA, err = ipa.NewFromString(clientIPAddr.String())
	if err != nil {
		srv.respondWithInternalServerError(rw)
		return nil, err
	}

	ipaRange, err = srv.iparc.GetRangeByIPAddress(clientIPA)
	if err != nil {
		srv.respondWithInternalServerError(rw)
		return nil, err
	}

	return ipaRange, nil
}

func (srv *Server) isIPARAllowed(
	ipaRange *ipar.IPAddressV4Range,
) (ok bool, countryCode string) {
	countryCode = ipaRange.GetCountry().Code()

	if countryCode == country.CodeUnknown {
		return srv.settings.AllowUnknownCountries, countryCode
	}

	_, countryCodeIsForbidden := srv.forbiddenCountryCodes[countryCode]

	return !countryCodeIsForbidden, countryCode
}
