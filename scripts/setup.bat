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
SET GitLink_V13=github.com/vault-thirteen
SET GitLink_SFRODB_Server=%GitLink_V13%/SFRODB/cmd/SFRODB/server@latest
SET GitLink_SFRODB_Client=%GitLink_V13%/SFRODB/cmd/SFRODB/client@latest
:: Note that 'SFHS' is now merged into the 'SFRODB' repository.
SET GitLink_SFHS_Server=%GitLink_V13%/SFRODB/cmd/SFHS/server@latest
SET GitLink_SPA_Server=%GitLink_V13%/SPA/cmd/spaServer@latest
SET GitLink_SPA_Proxy=%GitLink_V13%/SPA/cmd/proxy@latest
SET GitLink_SPA_Hasher=%GitLink_V13%/SPA/cmd/jsonHasher@latest
SET GitLink_SPA_Indexer=%GitLink_V13%/SPA/cmd/indexer@latest

SET V13_Folder=GOPATH\pkg\mod\github.com\vault-thirteen
SET TLS_CERT=cert\server-cert.pem
SET TLS_KEY=cert\server-key.pem
SET Main_Address=https://localhost

SET SFRODB_Common_Cache_Volume_Max=128000000
SET SFRODB_Icon_Cache_Volume_Max=8000000
SET SFRODB_Common_Item_Volume_Max=1000000
SET SFRODB_Icon_Item_Volume_Max=50000
SET SFRODB_Item_TTL=60
SET SFRODB_Base_Host=localhost
SET SFRODB_Base_Main_Port=2000
SET SFRODB_Base_Aux_Port=3000
SET SFRODB_Data_Folder=data

SET SFHS_Base_Host=localhost
SET SFHS_Base_Port=8000
SET SFHS_Base_Work_Mode=https
SET SFHS_Base_Certificate=..\SPA\Proxy\%TLS_CERT%
SET SFHS_Base_Key=..\SPA\Proxy\%TLS_KEY%
SET SFHS_Base_Db_Client_Pool_Size=10
SET SFHS_Base_TTL=300
SET SFHS_Base_CORS_Host=%Main_Address%

SET SPA_Host=%SFHS_Base_Host%
SET SPA_Port=%SFHS_Base_Port%
SET SPA_Base_Work_Mode=http
SET SPA_Base_Certificate=
SET SPA_Base_Key=
SET SPA_Base_TTL=%SFHS_Base_TTL%
SET SPA_Base_CORS_Host=%Main_Address%
SET SPA_Files=loader.js, styles.css, favicon.ico.png

SET SPA_Proxy_Host=0.0.0.0
SET SPA_Proxy_Port_Main=443
SET SPA_Proxy_Work_Mode=https
SET SPA_Proxy_Certificate=%TLS_CERT%
SET SPA_Proxy_Key=%TLS_KEY%
SET SPA_Proxy_TTL=%SFHS_Base_TTL%
SET SPA_Proxy_CORS_Host=%Main_Address%
SET SPA_Proxy_Target_Main=http://localhost:8000
SET SPA_Proxy_IPARC_DbFile=data\db\DB-1.csv.zip
SET SPA_Proxy_AllowUnknownCountries=yes
SET SPA_Proxy_ForbiddenCountryCodes=

SET SPA_Indexer_CategoryPaths=event, game, hard, life, media, motor, news, review, soft, tech
SET SPA_Indexer_ShouldCreateCategoryFolder=yes
SET SPA_Indexer_TopNewsCount=3
SET SPA_Indexer_MainServerAddress=%SFHS_Base_CORS_Host%:%SPA_Proxy_Port_Main%
SET SPA_Indexer_IconServerAddress=%SFHS_Base_CORS_Host%:%SPA_Proxy_Port_Icon%
SET SPA_Indexer_JpegServerAddress=%SFHS_Base_CORS_Host%:%SPA_Proxy_Port_Jpeg%
SET SPA_Indexer_JsonServerAddress=%SFHS_Base_CORS_Host%:%SPA_Proxy_Port_Json%

:: Part 0. Notes.
ECHO [91m[!] For some reason Windows Defender can sometimes think that an HTTP server, database server or a proxy server is a virus. Of course, this not true, and you can check all the source code of all the servers. If the antivirus decides that an HTTP server is a virus, open the list of latest scans and allow the latest quarantined compiled executable files, then re-use the script again. [!][0m

:: Part I. Get the executable files.
:Part1
SETLOCAL DisableDelayedExpansion
    ECHO Preparing Folders ...
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
	SET GOPATH=%CD%\GOPATH
	ECHO %GOPATH%
	
	:: SFRODB executable files.
	ECHO Installing SFRODB Server ...
	go install %GitLink_SFRODB_Server%
	IF %Errorlevel% NEQ 0 EXIT /b %Errorlevel%
	ECHO Installing SFRODB Client ...
	go install %GitLink_SFRODB_Client%
	IF %Errorlevel% NEQ 0 EXIT /b %Errorlevel%
	MOVE "GOPATH\bin\server.exe" "SFRODB\"
	MOVE "GOPATH\bin\client.exe" "SFRODB\"
	
	:: SFHS executable files.
	ECHO Installing SFHS Server ...
	go install %GitLink_SFHS_Server%
	IF %Errorlevel% NEQ 0 EXIT /b %Errorlevel%
	MOVE "GOPATH\bin\server.exe" "SFHS\"
	
	:: SPA executable files.
	ECHO Installing SPA Server ...
	go install %GitLink_SPA_Server%
	IF %Errorlevel% NEQ 0 EXIT /b %Errorlevel%
	RENAME "GOPATH\bin\spaServer.exe" "server.exe"
	MOVE "GOPATH\bin\server.exe" "SPA\Server\"
	::
	ECHO Installing SPA Proxy ...
	go install %GitLink_SPA_Proxy%
	IF %Errorlevel% NEQ 0 EXIT /b %Errorlevel%
	MOVE "GOPATH\bin\proxy.exe" "SPA\Proxy\"
	::
	ECHO Installing SPA Hasher ...
	go install %GitLink_SPA_Hasher%
	IF %Errorlevel% NEQ 0 EXIT /b %Errorlevel%
	RENAME "GOPATH\bin\jsonHasher.exe" "hasher.exe"
	MOVE "GOPATH\bin\hasher.exe" "SPA\Hasher\"
	::
	ECHO Installing SPA Indexer ...
	go install %GitLink_SPA_Indexer%
	IF %Errorlevel% NEQ 0 EXIT /b %Errorlevel%
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
	ECHO jpg %SFRODB_Icon_Cache_Volume_Max% %SFRODB_Icon_Item_Volume_Max% %SFRODB_Item_TTL%
) > "SFRODB\icon-db\settings.txt"
:: SFHS - IconDb.
MKDIR "SFHS\icon-db"
SET /A SFHS_Port=%SFHS_Base_Port% + %PortDelta%
(
	ECHO %SFHS_Base_Host%
	ECHO %SFHS_Port%
	ECHO %SFHS_Base_Work_Mode%
	ECHO %SFHS_Base_Certificate%
	ECHO %SFHS_Base_Key%
	ECHO %SFRODB_Base_Host%
	ECHO %DbMainPort%
	ECHO %DbAuxPort%
	ECHO %SFHS_Base_Db_Client_Pool_Size%
	ECHO jpg
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
	ECHO jpg %SFRODB_Common_Cache_Volume_Max% %SFRODB_Common_Item_Volume_Max% %SFRODB_Item_TTL%
) > "SFRODB\jpeg-db\settings.txt"
:: SFHS - JpegDb.
MKDIR "SFHS\jpeg-db"
SET /A SFHS_Port=%SFHS_Base_Port% + %PortDelta%
(
	ECHO %SFHS_Base_Host%
	ECHO %SFHS_Port%
	ECHO %SFHS_Base_Work_Mode%
	ECHO %SFHS_Base_Certificate%
	ECHO %SFHS_Base_Key%
	ECHO %SFRODB_Base_Host%
	ECHO %DbMainPort%
	ECHO %DbAuxPort%
	ECHO %SFHS_Base_Db_Client_Pool_Size%
	ECHO jpg
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
	ECHO json %SFRODB_Common_Cache_Volume_Max% %SFRODB_Common_Item_Volume_Max% %SFRODB_Item_TTL%
) > "SFRODB\json-db\settings.txt"
:: SFHS - JsonDb.
MKDIR "SFHS\json-db"
SET /A SFHS_Port=%SFHS_Base_Port% + %PortDelta%
(
	ECHO %SFHS_Base_Host%
	ECHO %SFHS_Port%
	ECHO %SFHS_Base_Work_Mode%
	ECHO %SFHS_Base_Certificate%
	ECHO %SFHS_Base_Key%
	ECHO %SFRODB_Base_Host%
	ECHO %DbMainPort%
	ECHO %DbAuxPort%
	ECHO %SFHS_Base_Db_Client_Pool_Size%
	ECHO json
	ECHO application/json
	ECHO %SFHS_Base_TTL%
	ECHO %SFHS_Base_CORS_Host%
) > "SFHS\json-db\settings.txt"

:: SPA Server.
(
	ECHO %SPA_Host%
	ECHO %SPA_Port%
	ECHO %SPA_Base_Work_Mode%
	IF "%SPA_Base_Certificate%" EQU "" ( ECHO: ) ELSE ( ECHO %SPA_Base_Certificate%)
	IF "%SPA_Base_Key%" EQU "" ( ECHO: ) ELSE ( ECHO %SPA_Base_Key%)
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

:: Proxy for SPA Server.
MKDIR "SPA\Proxy\main"
(
	ECHO %SPA_Proxy_Host%
	ECHO %SPA_Proxy_Port_Main%
	ECHO %SPA_Proxy_Work_Mode%
	ECHO %SPA_Proxy_Certificate%
	ECHO %SPA_Proxy_Key%
	ECHO %SPA_Proxy_TTL%
	ECHO:
	ECHO %SPA_Proxy_Target_Main%
	ECHO %SPA_Proxy_IPARC_DbFile%
	ECHO %SPA_Proxy_AllowUnknownCountries%
	IF "%SPA_Proxy_ForbiddenCountryCodes%" EQU "" ( ECHO: ) ELSE ( ECHO %SPA_Proxy_ForbiddenCountryCodes%)
	ECHO yes
) > "SPA\Proxy\main\settings.txt"

ECHO SUCCESSFUL SETUP
