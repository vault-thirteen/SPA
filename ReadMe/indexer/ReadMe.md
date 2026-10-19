# Indexer

A tool for creating index files for categories.

## Building
Use the `build.bat` script included with the source code.

## Installation
`go install github.com/vault-thirteen/SPA/cmd/indexer@latest`  

## Startup Parameters

### Server

`indexer.exe [Category] [SettingsFile]`  
`indexer.exe`  

Example:

`indexer.exe soft settings.txt`  
`indexer.exe news settings.txt`  
`indexer.exe tech`  

**Notes:**  
If the path to a configuration file is omitted, the default one is used.  
Default name of the configuration file is `settings.txt`.  

## Settings

Format of the settings' file for a server is quite simple. It uses line breaks 
as a separator between parameters. Further information can be found in the 
source code.
