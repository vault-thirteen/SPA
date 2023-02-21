# CRC32 JSON Hasher

This small tool is used for checking and writing CRC32 check sum of a JSON 
file.

This is not a universal tool, this tool works only with its specific data 
format. A separate executable file is used for checking and another one – for 
writing CRC32 check sums. A JavaScript function for the web browser to 
calculate CRC32 check sums is included.

## CRC

CRC means a cyclic redundancy check.

> A cyclic redundancy check (CRC) is an error-detecting code commonly used in 
digital networks and storage devices to detect accidental changes to digital 
data.

https://en.wikipedia.org/wiki/Cyclic_redundancy_check

This tool calculates a standard CRC-32 checksum of data using the IEEE 
polynomial.

## Why CRC ? 

Why not MD or SHA ?  
CRC is very fast and is still quite good for small amounts of data.  
CRC32 is very easy to be implemented in JavaScript on the web browser's side.

## Building

Use the `build.bat` script included with the source code.

## Installation

`go install github.com/vault-thirteen/SPA/cmd/jsonHasher@latest`   

## Usage

`jsonHasher.exe <action> <json-file>`

Example:  
`jsonHasher.exe hash "my.json"`  
`jsonHasher.exe check "my.json"`  
