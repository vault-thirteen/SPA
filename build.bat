:: This script builds the server.
@ECHO OFF

SET build_dir=_build_
SET exe_dir=cmd
SET spa_server_dir=spaServer
SET settings_file=settings.txt
SET assets_folder=assets
SET json_hasher_dir=jsonHasher
SET tools_folder="tools

MKDIR "%build_dir%"
MKDIR "%build_dir%\%spa_server_dir%"
MKDIR "%build_dir%\%tools_folder%"
MKDIR "%build_dir%\%tools_folder%\%json_hasher_dir%"

:: Build the SPA server.
CD "%exe_dir%\%spa_server_dir%"
go build
MOVE "%spa_server_dir%.exe" ".\..\..\%build_dir%\%spa_server_dir%\"
CD ".\..\..\"

:: Copy some additional files for the server.
COPY "%exe_dir%\%spa_server_dir%\%settings_file%" "%build_dir%\%spa_server_dir%\"
COPY "%exe_dir%\%spa_server_dir%\create-certificates.bat" "%build_dir%\%spa_server_dir%\"

:: Copy the assets.
COPY "%assets_folder%\favicon.ico.png" "%build_dir%\%spa_server_dir%\"
COPY "%assets_folder%\index.html" "%build_dir%\%spa_server_dir%\"
COPY "%assets_folder%\loader.js" "%build_dir%\%spa_server_dir%\"
COPY "%assets_folder%\styles.css" "%build_dir%\%spa_server_dir%\"

:: Build the CRC32 JSON Hasher.
CD "%exe_dir%\%json_hasher_dir%"
go build
MOVE "%json_hasher_dir%.exe" ".\..\..\%build_dir%\%tools_folder%\%json_hasher_dir%\"
CD ".\..\..\"
