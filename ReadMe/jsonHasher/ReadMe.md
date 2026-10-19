# CRC32 JSON Hasher

This small tool is used for checking and writing _CRC32_ check sum of a _JSON_ 
file.

This is not a universal tool, this tool works only with its specific data 
format. A separate executable file is used for checking and another one – for 
writing _CRC32_ check sums. A _JavaScript_ function for the web browser to 
calculate _CRC32_ check sums is included.

## CRC

_CRC_ means a cyclic redundancy check.

> A cyclic redundancy check (CRC) is an error-detecting code commonly used in 
digital networks and storage devices to detect accidental changes to digital 
data.

https://en.wikipedia.org/wiki/Cyclic_redundancy_check

This tool calculates a standard _CRC-32_ checksum of data using the _IEEE_ 
polynomial.

## Why CRC ? 

Why not _MD_ or _SHA_ ?  
_CRC_ is very fast and is still quite good for small amounts of data.  
_CRC32_ is very easy to be implemented in _JavaScript_ on the web browser's side.

## Building

Use the `build.bat` script included with the source code.

## Installation

`go install github.com/vault-thirteen/SPA/cmd/jsonHasher@latest`   

## Usage

`jsonHasher.exe <action> <json-file>`

Example:  
`jsonHasher.exe hash "my.json"`  
`jsonHasher.exe check "my.json"`  
