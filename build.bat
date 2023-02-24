:: This script builds the server.
@ECHO OFF

SET build_dir=_build_
SET exe_dir=cmd
SET spa_server_dir=spaServer
SET settings_file=settings.txt
SET assets_folder=assets
SET tools_folder="tools
SET json_hasher_dir=jsonHasher
SET indexer_dir=indexer
SET spa_rp_server_dir=proxy

:: Create the folders.
MKDIR "%build_dir%"
MKDIR "%build_dir%\%spa_server_dir%"
MKDIR "%build_dir%\%tools_folder%"
MKDIR "%build_dir%\%tools_folder%\%json_hasher_dir%"
MKDIR "%build_dir%\%tools_folder%\%indexer_dir%"
MKDIR "%build_dir%\%spa_rp_server_dir%"

:: Build the SPA Server.
CD "%exe_dir%\%spa_server_dir%"
go build
MOVE "%spa_server_dir%.exe" ".\..\..\%build_dir%\%spa_server_dir%\"
CD ".\..\..\"

:: Copy some additional files for the server.
COPY "%exe_dir%\%spa_server_dir%\%settings_file%" "%build_dir%\%spa_server_dir%\"

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

:: Build the Indexer.
CD "%exe_dir%\%indexer_dir%"
go build
MOVE "%indexer_dir%.exe" ".\..\..\%build_dir%\%tools_folder%\%indexer_dir%\"
CD ".\..\..\"

:: Copy some additional files for the indexer.
COPY "%exe_dir%\%indexer_dir%\%settings_file%" "%build_dir%\%tools_folder%\%indexer_dir%\"

:: Build the SPA Reverse Proxy Server.
CD "%exe_dir%\%spa_rp_server_dir%"
go build
MOVE "%spa_rp_server_dir%.exe" ".\..\..\%build_dir%\%spa_rp_server_dir%\"
CD ".\..\..\"

:: Copy some additional files for the server.
COPY "%exe_dir%\%spa_rp_server_dir%\%settings_file%" "%build_dir%\%spa_rp_server_dir%\"
