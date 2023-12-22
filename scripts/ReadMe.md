# Scripts

This folder contains installation scripts for the _SPA_ bundle.

These scripts must be used outside the repository folder, in a folder where 
you wish to install the _SPA_ bundle. This means that you must first copy these 
scripts to a folder where you plan to install the _SPA_ bundle, and then run them 
in that separate folder. 

These scripts do not use the source folder while they fetch all data from the 
Internet.

The order of scripts to use is following.

1. `setup.bat`
2. `create-certificates.bat`
3. `move-certificates.bat`
4. `start.bat`
5. `stop.bat`

If you already have certificates in a correct folder, you can skip the steps 2 
and 3. If you use the starter script without prior creation of certificates, 
it will throw some strange errors because _SFHS_ server will be unable to 
start. Due to complexity of _Batch_ scripts in _Windows_ O.S., not all errors 
are checked in the starter script, so be careful with the execution order.

In order to use the `create-certificates.bat` script, you need to have 
_OpenSSL_ installed and reachable in your `PATH` environmental variable.
