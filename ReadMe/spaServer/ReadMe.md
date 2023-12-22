# SPA Server

An _HTTP_ server for a single page application in _JavaScript_.

An example of an index _HTML_ page, a loader in _JavaScript_ and a _CSS_ style 
sheet are included. A favicon file using the _PNG_ format is also included for 
reference.

## Building
Use the `build.bat` script included with the source code.

## Installation
`go install github.com/vault-thirteen/SPA/cmd/spaServer@latest`  

## Startup Parameters

### Server

`spaServer.exe <path-to-configuration-file>`  
`spaServer.exe`  

Example:

`spaServer.exe "settings.txt"`  
`spaServer.exe`  

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
8. List of the cached files related to the _JavaScript_ single page 
application, except the index HTML page which is hardcoded – a _JavaScript_ 
loader script, a _CSS_ style sheet and a favicon file.

**Notes:**  
This server is an SPA server. This means that it does not serve ordinary 
data files having the contents of the whole website. It serves only those files 
which are required to start the _JavaScript_ SPA itself and no more.

## Links Format

Links in the _JSON_ data files must not start with a slash symbol.  
Slash symbol is automatically inserted by the _JavaScript_ router.  
The same rule works for icons.
