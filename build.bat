:: This script builds the server.
@ECHO OFF

SET build_dir=_build_
SET exe_dir=cmd
SET server_dir=server
SET settings_file=settings.dat
SET server_starter_script=start-server.bat

MKDIR "%build_dir%"

:: Build the server.
CD "%exe_dir%\%server_dir%"
go build
MOVE "%server_dir%.exe" ".\..\..\%build_dir%\"
CD ".\..\..\"

:: Copy some additional files for the server.
COPY "%exe_dir%\%server_dir%\%settings_file%" "%build_dir%\"
COPY "%exe_dir%\%server_dir%\%server_starter_script%" "%build_dir%\"
