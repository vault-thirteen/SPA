# Proxy

A Reverse Proxy Server for the _SPA_ Server.

## Building
Use the `build.bat` script included with the source code.

## Installation
`go install github.com/vault-thirteen/SPA/cmd/proxy@latest`  

## Startup Parameters

### Server

`proxy.exe <path-to-configuration-file>`  
`proxy.exe`  

Example:

`proxy.exe "settings.txt"`  
`proxy.exe`  

**Notes:**  
If the path to a configuration file is omitted, the default one is used.  
Default name of the configuration file is `settings.txt`.  

## Settings

Format of the settings' file for a server is quite simple. It uses line breaks 
as a separator between parameters. Described below are meanings of each line.

1. Server's hostname.
2. Server's listen port.
3. Work mode: _HTTP_ or _HTTPS_.
4. Path to the certificate file for the _HTTPS_ work mode.
5. Path to the key file for the _HTTPS_ work mode.
6. TTL of served files, i.e. value of the `max-age` field of the
   `Cache-Control` _HTTP_ header.
7. Allowed origin for _HTTP_ CORS, i.e. value of the
   `Access-Control-Allow-Origin` _HTTP_ header.
8. Address of the target server.
9. Path to the _IPARC_ database file.
10. Boolean flag showing that unknown countries should be allowed.
11. Comma separated list of two-letter codes of forbidden countries.
12. Boolean flag showing that server is the main proxy server.

## HTTP Headers
* The built-in reverse proxy server automatically sets the `X-Forwarded-For` 
_HTTP_ header to contain the client's IP address.
* Client's country code is written to the `X-ClientCountryCode` _HTTP_ header.
