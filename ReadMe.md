# SPA

An _HTTP_ server for a single page application in _JavaScript_ and related 
tools.

For more information about each component see each component's ReadMe file 
located in an appropriate sub-folder of the `ReadMe` folder.

## Components
* SPA Server
* An example of a page loader and router written in JavaScript
* Complementary reference files (_CSS_ style sheet, etc.) 
* CRC32 hash tool for _JSON_ files
* Indexer
* Reverse Proxy Server
* Other components (external)

### Components' ReadMe Files
* SPA Server – [ReadMe](ReadMe/spaServer/ReadMe.md)
* CRC32 hash tool for JSON files – [ReadMe](ReadMe/jsonHasher/ReadMe.md)
* Indexer – [ReadMe](ReadMe/indexer/ReadMe.md)
* Reverse Proxy Server – [ReadMe](ReadMe/proxy/ReadMe.md)
* External components:
  * _SFRODB_ database – [ReadMe](https://github.com/vault-thirteen/SFRODB)
  * _SFHS_ interface – [ReadMe](https://github.com/vault-thirteen/SFHS)
  * _IPARC_ library – [ReadMe](https://github.com/vault-thirteen/IPARC)

## Structural Diagram
![Structural Diagram](https://github.com/vault-thirteen/SPA/blob/839a7b32913de19863bac548de8167c6e5298909/Documentation/SPA%20Structural%20Diagram.png)

## Installation
1. Create a new separate folder which will be used as a root folder for your 
installation.


2. Copy following scripts from the `scripts` folder into this root folder:
   - `setup.bat`
   - `start.bat`
   - `stop.bat`


3. Modify the setup script to your needs:
   - Set the domain name in the `Main_Address` variable;
   - Tweak cache settings using following variables if you do not like the 
default values:
     - `SFRODB_Common_Cache_Volume_Max` or `SFRODB_Icon_Cache_Volume_Max`
     - `SFRODB_Common_Item_Volume_Max` or `SFRODB_Icon_Item_Volume_Max`
     - `SFRODB_Item_TTL`
   - Customize the indexer settings:
     - `SPA_Indexer_TopNewsCount`
   - All other settings should not be touched as they are there for a reason.


4. Run the setup script and ensure than no errors occur. It may happen that _Go_ 
language changes in future and this script stops working. The reasons are 
stated inside the setup script as comments.


5. Either copy your _SSL/TLS_ certificates into the `SPA\Proxy\cert` folder if you 
already have them, or run the `create-certificates.bat` script located at 
`SPA\Proxy` which will automatically create and place your self-signed 
certificates into the `cert` folder.


6. Fill all the data folders with your own content. 
It is not included with this bundle, as it is only a framework.
Available data folders:
    - `SFRODB\icon-db\data`
    - `SFRODB\jpeg-db\data`
    - `SFRODB\json-db\data`


7. Run the `start.bat` script to start all the servers.


8. Run the `stop.bat` script to stop all the servers.


**Notes**

* If you would like to use this product on a distributed platform (i.e. on 
separate network machines), all the configurations and the setup script should 
be changed accordingly.


* Server addresses, including host names and port numbers, and numerous other 
settings are hard-coded not only in the `setup.bat` script, but also in the 
front-end loader script (`loader.js`), a start web page (`index.html`) and some 
other places. So, do not touch anything unless you are sure about what 
you are doing.


* This product serves static content which must be re-indexed after each change. 
This means that after each change of the content in your data folders, a new 
indexation must be run for each category of the modified content. This is the 
key difference between this framework and most other website engines available 
in the network. Also note that each re-indexation must be followed by a cache 
cleaning. For more information see the documentation of external components. 
