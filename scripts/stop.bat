:: Stop script for the SPA bundle ::

@ECHO OFF

TASKKILL /FI "WINDOWTITLE eq SPA-Proxy-Main"

TASKKILL /FI "WINDOWTITLE eq SPA-Server"
TASKKILL /FI "WINDOWTITLE eq SFHS-IconDb"
TASKKILL /FI "WINDOWTITLE eq SFHS-JpegDb"
TASKKILL /FI "WINDOWTITLE eq SFHS-JsonDb"

TASKKILL /FI "WINDOWTITLE eq SFRODB-IconDb"
TASKKILL /FI "WINDOWTITLE eq SFRODB-JpegDb"
TASKKILL /FI "WINDOWTITLE eq SFRODB-JsonDb"
