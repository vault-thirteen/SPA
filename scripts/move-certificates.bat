SET InputFolder=cert
SET OutputFolder=SPA\Proxy\cert

MOVE %InputFolder%\ca-cert.pem %OutputFolder%\
MOVE %InputFolder%\ca-key.pem %OutputFolder%\
MOVE %InputFolder%\server-key.pem %OutputFolder%\
MOVE %InputFolder%\server-cert.pem %OutputFolder%\
MOVE %InputFolder%\server-req.pem %OutputFolder%\

RMDIR %InputFolder%
