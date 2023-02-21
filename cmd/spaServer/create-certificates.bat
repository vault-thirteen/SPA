SET openssl=openssl

:: Create private key for CA.
%openssl% genrsa 8192 > ca-key.pem

:: Create X509 certificate for CA.
%openssl% req -new -x509 -nodes -days 365000 -key ca-key.pem -out ca-cert.pem

:: Create certificate and key for server.
%openssl% req -newkey rsa:8192 -nodes -days 365000 -keyout server-key.pem -out server-req.pem
%openssl% x509 -req -days 365000 -set_serial 01 -in server-req.pem -out server-cert.pem -CA ca-cert.pem -CAkey ca-key.pem

:: Create certificate and key for client.
::%openssl% req -newkey rsa:8192 -nodes -days 365000 -keyout client-key.pem -out client-req.pem
::%openssl% x509 -req -days 365000 -set_serial 01 -in client-req.pem -out client-cert.pem -CA ca-cert.pem -CAkey ca-key.pem

:: Verification.
%openssl% verify -CAfile ca-cert.pem ca-cert.pem server-cert.pem
%openssl% verify -CAfile ca-cert.pem ca-cert.pem client-cert.pem
