:: Installation script for the SPA bundle ::

:: Notes.
::
:: Due to an old bug (they call it a feature, of course) in Go programming
:: language, it is not possible to use the 'go install' command with a custom
:: installation path like it is possible with 'go build' command. What does it
:: mean ? It is possible to build an executable file into a custom folder using
:: the '-o' command line argument, but it is still not possible to use '-o'
:: with 'go install' command. Because of all this shame, we have to re-invent
:: the wheel and write such stupid scripts like this one.

@ECHO OFF

:: Settings.
SET TLS_CERT=cert\server-cert.pem
SET TLS_KEY=cert\server-key.pem
SET Main_Address=https://localhost

SET SFRODB_Cache_Volume_Max=128000000
SET SFRODB_Item_Volume_Max=1000000
SET SFRODB_Item_TTL=60
SET SFRODB_Base_Host=localhost
SET SFRODB_Base_Main_Port=2000
SET SFRODB_Base_Aux_Port=3000
SET SFRODB_Data_Folder=data

SET SFHS_Base_Host=localhost
SET SFHS_Base_Port=8000
SET SFHS_Base_Work_Mode=http
SET SFHS_Base_Certificate=none
SET SFHS_Base_Key=none
SET SFHS_Base_Db_Client_Pool_Size=10
SET SFHS_Base_TTL=300
SET SFHS_Base_CORS_Host=%Main_Address%

SET SPA_Host=%SFHS_Base_Host%
SET SPA_Port=%SFHS_Base_Port%
SET SPA_Base_Work_Mode=%SFHS_Base_Work_Mode%
SET SPA_Base_Certificate=none
SET SPA_Base_Key=none
SET SPA_Base_TTL=%SFHS_Base_TTL%
SET SPA_Base_CORS_Host=http://localhost
SET SPA_Files=loader.js, styles.css, favicon.ico.png

SET SPA_Proxy_Host=0.0.0.0
SET SPA_Proxy_Port_Main=443
SET SPA_Proxy_Port_Icon=1101
SET SPA_Proxy_Port_Jpeg=1102
SET SPA_Proxy_Port_Json=1103
SET SPA_Proxy_Work_Mode=https
SET SPA_Proxy_Certificate=%TLS_CERT%
SET SPA_Proxy_Key=%TLS_KEY%
SET SPA_Proxy_TTL=%SFHS_Base_TTL%
SET SPA_Proxy_CORS_Host=
SET SPA_Proxy_Target_Main=http://localhost:8000
SET SPA_Proxy_Target_Icon=http://localhost:8001
SET SPA_Proxy_Target_Jpeg=http://localhost:8002
SET SPA_Proxy_Target_Json=http://localhost:8003
SET SPA_Proxy_IPARC_DbFile=data\db\DB-1.csv.zip

SET SPA_Indexer_CategoryPaths=event, game, hard, life, media, motor, news, review, soft, tech
SET SPA_Indexer_ShouldCreateCategoryFolder=yes
SET SPA_Indexer_TopNewsCount=3
SET SPA_Indexer_MainServerAddress=%SFHS_Base_CORS_Host%:%SPA_Proxy_Port_Main%
SET SPA_Indexer_IconServerAddress=%SFHS_Base_CORS_Host%:%SPA_Proxy_Port_Icon%
SET SPA_Indexer_JpegServerAddress=%SFHS_Base_CORS_Host%:%SPA_Proxy_Port_Jpeg%
SET SPA_Indexer_JsonServerAddress=%SFHS_Base_CORS_Host%:%SPA_Proxy_Port_Json%

SET V13_Folder=GOPATH\pkg\mod\github.com\vault-thirteen

::GOTO Part1
::GOTO Part2

:: Part I. Get the executable files.
:Part1
SETLOCAL DisableDelayedExpansion
	MKDIR "SFHS"
	MKDIR "SFRODB"
	MKDIR "SPA"
	MKDIR "SPA\Server"
	MKDIR "SPA\Proxy"
	MKDIR "SPA\Proxy\cert"
	MKDIR "SPA\Proxy\data"
	MKDIR "SPA\Proxy\data\db"
	MKDIR "SPA\Hasher"
	MKDIR "SPA\Indexer"
	
	ECHO %GOPATH%
	MKDIR GOPATH
	SET GOPATH="%CD%\GOPATH"
	ECHO %GOPATH%
	
	:: SFRODB executable files.
	go install github.com/vault-thirteen/SFRODB/cmd/server@latest
	go install github.com/vault-thirteen/SFRODB/cmd/client@latest
	MOVE "GOPATH\bin\server.exe" "SFRODB\"
	MOVE "GOPATH\bin\client.exe" "SFRODB\"
	
	:: SFHS executable files.
	go install github.com/vault-thirteen/SFHS/cmd/server@latest
	MOVE "GOPATH\bin\server.exe" "SFHS\"
	
	:: SPA executable files.
	go install github.com/vault-thirteen/SPA/cmd/spaServer@latest
	RENAME "GOPATH\bin\spaServer.exe" "server.exe"
	MOVE "GOPATH\bin\server.exe" "SPA\Server\"
	::
	go install github.com/vault-thirteen/SPA/cmd/proxy@latest
	MOVE "GOPATH\bin\proxy.exe" "SPA\Proxy\"
	::
	go install github.com/vault-thirteen/SPA/cmd/jsonHasher@latest
	RENAME "GOPATH\bin\jsonHasher.exe" "hasher.exe"
	MOVE "GOPATH\bin\hasher.exe" "SPA\Hasher\"
	::
	go install github.com/vault-thirteen/SPA/cmd/indexer@latest
	MOVE "GOPATH\bin\indexer.exe" "SPA\Indexer\"
	
	:: Looks like the stupidity was not in the Git. The stupidity is in 
	:: the Go language itself. Issue #26456, opened on 2018-07-19.
	:: https://github.com/golang/go/issues/26456. Almost 11 years have 
	:: passed since the first release of the Go programming language in 
	:: the year 2012, or 2012-03-28 according to the release history
	:: page located at https://go.dev/doc/devel/release. You wanted to 
	:: create a replacement for the good old C language. And what 
	:: happened instead ? ...
	
	
	:: Try to find the SPA source folder in the GOPATH.
	CD "%V13_Folder%"
	
	:: Count the folders having a name starting with "!s!p!a".
	SET Folder_Pattern=^^^^!s^^^^!p^^^^!a*
	SETLOCAL EnableDelayedExpansion
	DIR %Folder_Pattern% /A:D /B
	SET c=0
	FOR /F %%i IN ('DIR %Folder_Pattern% /A:D /B') DO (	SET /A c=!c! + 1 )
	ECHO Total number of folders found is !c!.
	IF !c! NEQ 1 ( 
		ECHO Must be a single folder for the SPA.
		CD "..\..\..\..\..\"
		EXIT /B 1 )
	ENDLOCAL
	
	:: Get the single folder having a name starting with "!s!p!a".
	SET Folder_Pattern=!s!p!a*
	FOR /F %%i IN ('DIR %Folder_Pattern% /A:D /B') DO (
		SET Folder_Name=%%i
		ECHO Found a folder: %%i )
	ECHO %Folder_Name%
	
	:: Open required folders.
	CD "%Folder_Name%\assets"
	CD "..\..\"
	CD "%Folder_Name%\scripts"
	CD "..\..\"
	CD "..\..\..\..\..\"

	:: Copy the SPA assets.
	COPY "%V13_Folder%\%Folder_Name%\assets\*" "SPA\Server\"

	:: Copy SPA scripts.
	COPY "%V13_Folder%\%Folder_Name%\scripts\create-certificates.bat" "SPA\Proxy\"


	:: Try to find the IPARC source folder in the GOPATH.
	CD "%V13_Folder%"

	:: Count the folders having a name starting with "!i!p!a!r!c".
	SET Folder_Pattern=^^^^!i^^^^!p^^^^!a^^^^!r^^^^!c*
	SETLOCAL EnableDelayedExpansion
	DIR %Folder_Pattern% /A:D /B
	SET c=0
	FOR /F %%i IN ('DIR %Folder_Pattern% /A:D /B') DO (	SET /A c=!c! + 1 )
	ECHO Total number of folders found is !c!.
	IF !c! NEQ 1 (
		ECHO Must be a single folder for the IPARC.
		CD "..\..\..\..\..\"
		EXIT /B 1 )
	ENDLOCAL

	:: Get the single folder having a name starting with "!i!p!a!r!c".
	SET Folder_Pattern=!i!p!a!r!c*
	FOR /F %%i IN ('DIR %Folder_Pattern% /A:D /B') DO (
		SET Folder_Name=%%i
		ECHO Found a folder: %%i )
	ECHO %Folder_Name%

	:: Open required folders.
	CD "%Folder_Name%\data\db"
	CD "..\..\..\"
	CD "..\..\..\..\..\"

	:: Copy the IPARC database.
	COPY "%V13_Folder%\%Folder_Name%\data\db\*" "SPA\Proxy\data\db\"
	
	:: Clear the temporary folder.
	RMDIR /S /Q "GOPATH"
	
ENDLOCAL

:: Part II. Prepare configuration files.
:Part2

:: SFRODB - IconDb.
MKDIR "SFRODB\icon-db"
MKDIR "SFRODB\icon-db\%SFRODB_Data_Folder%"
SET PortDelta=1
SET /A DbMainPort=%SFRODB_Base_Main_Port% + %PortDelta%
SET /A DbAuxPort=%SFRODB_Base_Aux_Port% + %PortDelta%
(
	ECHO %SFRODB_Base_Host%
	ECHO %DbMainPort%
	ECHO %DbAuxPort%
	ECHO icon-db\%SFRODB_Data_Folder%
	ECHO .txt %SFRODB_Cache_Volume_Max% %SFRODB_Item_Volume_Max% %SFRODB_Item_TTL%
	ECHO icon-db\%SFRODB_Data_Folder%
	ECHO .jpg %SFRODB_Cache_Volume_Max% %SFRODB_Item_Volume_Max% %SFRODB_Item_TTL%
) > "SFRODB\icon-db\settings.txt"
:: SFHS - IconDb.
MKDIR "SFHS\icon-db"
SET /A SFHS_MainPort=%SFHS_Base_Port% + %PortDelta%
(
	ECHO %SFHS_Base_Host%
	ECHO %SFHS_MainPort%
	ECHO %SFHS_Base_Work_Mode%
	ECHO %SFHS_Base_Certificate%
	ECHO %SFHS_Base_Key%
	ECHO %SFRODB_Base_Host%
	ECHO %DbMainPort%
	ECHO %DbAuxPort%
	ECHO %SFHS_Base_Db_Client_Pool_Size%
	ECHO .jpg
	ECHO image/jpeg
	ECHO %SFHS_Base_TTL%
	ECHO %SFHS_Base_CORS_Host%
) > "SFHS\icon-db\settings.txt"

:: SFRODB - JpegDb.
MKDIR "SFRODB\jpeg-db"
MKDIR "SFRODB\jpeg-db\%SFRODB_Data_Folder%"
SET PortDelta=2
SET /A DbMainPort=%SFRODB_Base_Main_Port% + %PortDelta%
SET /A DbAuxPort=%SFRODB_Base_Aux_Port% + %PortDelta%
(
	ECHO %SFRODB_Base_Host%
	ECHO %DbMainPort%
	ECHO %DbAuxPort%
	ECHO jpeg-db\%SFRODB_Data_Folder%
	ECHO .txt %SFRODB_Cache_Volume_Max% %SFRODB_Item_Volume_Max% %SFRODB_Item_TTL%
	ECHO jpeg-db\%SFRODB_Data_Folder%
	ECHO .jpg %SFRODB_Cache_Volume_Max% %SFRODB_Item_Volume_Max% %SFRODB_Item_TTL%
) > "SFRODB\jpeg-db\settings.txt"
:: SFHS - JpegDb.
MKDIR "SFHS\jpeg-db"
SET /A SFHS_MainPort=%SFHS_Base_Port% + %PortDelta%
(
	ECHO %SFHS_Base_Host%
	ECHO %SFHS_MainPort%
	ECHO %SFHS_Base_Work_Mode%
	ECHO %SFHS_Base_Certificate%
	ECHO %SFHS_Base_Key%
	ECHO %SFRODB_Base_Host%
	ECHO %DbMainPort%
	ECHO %DbAuxPort%
	ECHO %SFHS_Base_Db_Client_Pool_Size%
	ECHO .jpg
	ECHO image/jpeg
	ECHO %SFHS_Base_TTL%
	ECHO %SFHS_Base_CORS_Host%
) > "SFHS\jpeg-db\settings.txt"

:: SFRODB - JsonDb.
MKDIR "SFRODB\json-db"
MKDIR "SFRODB\json-db\%SFRODB_Data_Folder%"
SET PortDelta=3
SET /A DbMainPort=%SFRODB_Base_Main_Port% + %PortDelta%
SET /A DbAuxPort=%SFRODB_Base_Aux_Port% + %PortDelta%
(
	ECHO %SFRODB_Base_Host%
	ECHO %DbMainPort%
	ECHO %DbAuxPort%
	ECHO json-db\%SFRODB_Data_Folder%
	ECHO .txt %SFRODB_Cache_Volume_Max% %SFRODB_Item_Volume_Max% %SFRODB_Item_TTL%
	ECHO json-db\%SFRODB_Data_Folder%
	ECHO .json %SFRODB_Cache_Volume_Max% %SFRODB_Item_Volume_Max% %SFRODB_Item_TTL%
) > "SFRODB\json-db\settings.txt"
:: SFHS - JsonDb.
MKDIR "SFHS\json-db"
SET /A SFHS_MainPort=%SFHS_Base_Port% + %PortDelta%
(
	ECHO %SFHS_Base_Host%
	ECHO %SFHS_MainPort%
	ECHO %SFHS_Base_Work_Mode%
	ECHO %SFHS_Base_Certificate%
	ECHO %SFHS_Base_Key%
	ECHO %SFRODB_Base_Host%
	ECHO %DbMainPort%
	ECHO %DbAuxPort%
	ECHO %SFHS_Base_Db_Client_Pool_Size%
	ECHO .json
	ECHO application/json
	ECHO %SFHS_Base_TTL%
	ECHO %SFHS_Base_CORS_Host%
) > "SFHS\json-db\settings.txt"

:: SPA Server.
(
	ECHO %SPA_Host%
	ECHO %SPA_Port%
	ECHO %SPA_Base_Work_Mode%
	ECHO %SPA_Base_Certificate%
	ECHO %SPA_Base_Key%
	ECHO %SPA_Base_TTL%
	ECHO %SPA_Base_CORS_Host%
	ECHO %SPA_Files%
) > "SPA\Server\settings.txt"

:: SPA Indexer.
(
	ECHO %SFRODB_Data_Folder%
	ECHO %SPA_Indexer_CategoryPaths%
	ECHO %SPA_Indexer_ShouldCreateCategoryFolder%
	ECHO %SPA_Indexer_TopNewsCount%
	ECHO %SPA_Indexer_MainServerAddress%
	ECHO %SPA_Indexer_JsonServerAddress%
	ECHO %SPA_Indexer_IconServerAddress%
	ECHO %SPA_Indexer_JpegServerAddress%
) > "SPA\Indexer\settings.txt"

:: SPA Proxy for all resource servers.

:: 1. SPA Proxy for Main Server.
MKDIR "SPA\Proxy\main"
(
	ECHO %SPA_Proxy_Host%
	ECHO %SPA_Proxy_Port_Main%
	ECHO %SPA_Proxy_Work_Mode%
	ECHO %SPA_Proxy_Certificate%
	ECHO %SPA_Proxy_Key%
	ECHO %SPA_Proxy_TTL%
	IF [%SPA_Proxy_CORS_Host%]==[] ( ECHO: ) ELSE ( ECHO %SPA_Proxy_CORS_Host% )
	ECHO %SPA_Proxy_Target_Main%
	ECHO %SPA_Proxy_IPARC_DbFile%
) > "SPA\Proxy\main\settings.txt"

:: 2. SPA Proxy for IconDb.
MKDIR "SPA\Proxy\icon-db"
(
	ECHO %SPA_Proxy_Host%
	ECHO %SPA_Proxy_Port_Icon%
	ECHO %SPA_Proxy_Work_Mode%
	ECHO %SPA_Proxy_Certificate%
	ECHO %SPA_Proxy_Key%
	ECHO %SPA_Proxy_TTL%
	IF [%SPA_Proxy_CORS_Host%]==[] ( ECHO: ) ELSE ( ECHO %SPA_Proxy_CORS_Host% )
	ECHO %SPA_Proxy_Target_Icon%
	ECHO %SPA_Proxy_IPARC_DbFile%
) > "SPA\Proxy\icon-db\settings.txt"

:: 3. SPA Proxy for JpegDb.
MKDIR "SPA\Proxy\jpeg-db"
(
	ECHO %SPA_Proxy_Host%
	ECHO %SPA_Proxy_Port_Jpeg%
	ECHO %SPA_Proxy_Work_Mode%
	ECHO %SPA_Proxy_Certificate%
	ECHO %SPA_Proxy_Key%
	ECHO %SPA_Proxy_TTL%
	IF [%SPA_Proxy_CORS_Host%]==[] ( ECHO: ) ELSE ( ECHO %SPA_Proxy_CORS_Host% )
	ECHO %SPA_Proxy_Target_Jpeg%
	ECHO %SPA_Proxy_IPARC_DbFile%
) > "SPA\Proxy\jpeg-db\settings.txt"

:: 4. SPA Proxy for JsonDb.
MKDIR "SPA\Proxy\json-db"
(
	ECHO %SPA_Proxy_Host%
	ECHO %SPA_Proxy_Port_Json%
	ECHO %SPA_Proxy_Work_Mode%
	ECHO %SPA_Proxy_Certificate%
	ECHO %SPA_Proxy_Key%
	ECHO %SPA_Proxy_TTL%
	IF [%SPA_Proxy_CORS_Host%]==[] ( ECHO: ) ELSE ( ECHO %SPA_Proxy_CORS_Host% )
	ECHO %SPA_Proxy_Target_Json%
	ECHO %SPA_Proxy_IPARC_DbFile%
) > "SPA\Proxy\json-db\settings.txt"

