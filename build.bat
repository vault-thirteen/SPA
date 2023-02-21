:: This script builds the server.
@ECHO OFF

SET build_dir=_build_
SET exe_dir=cmd
SET server_dir=spaServer
SET settings_file=settings.txt
SET assets_folder=assets

MKDIR "%build_dir%"

:: Build the server.
CD "%exe_dir%\%server_dir%"
go build
MOVE "%server_dir%.exe" ".\..\..\%build_dir%\"
CD ".\..\..\"

:: Copy some additional files for the server.
COPY "%exe_dir%\%server_dir%\%settings_file%" "%build_dir%\"

:: Copy the assets.
COPY "%assets_folder%\favicon.ico.png" "%build_dir%\"
COPY "%assets_folder%\index.html" "%build_dir%\"
COPY "%assets_folder%\loader.js" "%build_dir%\"
COPY "%assets_folder%\styles.css" "%build_dir%\"
