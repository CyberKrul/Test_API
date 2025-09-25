## testAPI
An attempt to create a full fledged API, to illustrate the layered model, with dependency injection, and decoupled interactions.
Will be utilizing the *Gin* Framework for HTTP and JSON handling.

### API interactions
There are three* interactions planned:
1. Insertion: inserting **Serial Number** and **FirmwareVersion**
2. UpdateMesh: updating **Mesh Config** to True
3. UpdateMesh: updating **Mesh Config** to False

### Internal Data Model
This is the basic data model as it will be stored in the database (hidden from APIs)
|SNo|firmware_version|current_firmware_version|mesh_config|app_config|kc_config|
|int|int             |bool                     |  bool     |  bool    |  bool   |

### Features
Here are the core features to be implemented in each layer:
1. Model Layer:
    * Will hold the internal model of the data
    * Will create a constructor to initialize it
2. Repository Layer:
    * Will hold two files:
        1. the terms of the API contract 
            * will define an interface that holds the terms of the contract
            * An initializer for the contract
        2. the database interactions using a general *sql.DB 
            * all interactions will be defined as operating on an internal data object
3. Service Layer:
    * will define the business logic- data validation, insertion and update conditions and it will have its own 
    structure and constructor - to consume the interface
4. Handler Layer:
    * This will handle all things HTTP and JSON related, interacting with the service layer below it
5. Main.go:
    * This file will be the creation of the HTTP server, pulling in database credentials from a .env file and creating a connection to the database. It will make calls to the service layer to process the different routes of the API.