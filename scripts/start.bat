:: Start script for the SPA bundle ::

@ECHO OFF

:: Settings.
SET SettingsFile=settings.txt
SET SFRODB_Folder=SFRODB
SET SFHS_Folder=SFHS
SET SPA_Folder=SPA

CD %SFRODB_Folder%
START "SFRODB-IconDb" "server.exe" "icon-db\%SettingsFile%"
START "SFRODB-JpegDb" "server.exe" "jpeg-db\%SettingsFile%"
START "SFRODB-JsonDb" "server.exe" "json-db\%SettingsFile%"
CD ..

CD %SPA_Folder%\Server
START "SPA-Server" "server.exe" "%SettingsFile%"
CD ..\..

CD %SFHS_Folder%
START "SFHS-IconDb" "server.exe" "icon-db\%SettingsFile%"
START "SFHS-JpegDb" "server.exe" "jpeg-db\%SettingsFile%"
START "SFHS-JsonDb" "server.exe" "json-db\%SettingsFile%"
CD ..

CD %SPA_Folder%\Proxy
SET SleepTimeout=3
START "SPA-Proxy-IconDb" "proxy.exe" "icon-db\%SettingsFile%"
TIMEOUT /T %SleepTimeout% /nobreak
START "SPA-Proxy-JpegDb" "proxy.exe" "jpeg-db\%SettingsFile%"
TIMEOUT /T %SleepTimeout% /nobreak
START "SPA-Proxy-JsonDb" "proxy.exe" "json-db\%SettingsFile%"
TIMEOUT /T %SleepTimeout% /nobreak
START "SPA-Proxy-Main" "proxy.exe" "main\%SettingsFile%"
TIMEOUT /T %SleepTimeout% /nobreak
CD ..\..
