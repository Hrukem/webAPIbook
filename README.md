# webAPIbook

An application implementing a web API for storing book records in a database.

Implemented requests:  
GET /books  
returns a list of all objects in response to the request

GET /books?id=xxx  
in response to the request, it returns a record from the database at the specified id

PUT /book  
creates a new entry in the database returns the id of the created object

POST /book?id=xxx  
updates a record in the database by the specified id

DELETE /book?id=xxx  
deletes an entry by the specified id from the database

To run a project on a machine, docker-compose must be installed on it. To run on a local machine, clone the repository
with the command

- git clone https://github.com/Hrukem/webAPIbook

Go to the root folder of the project, where the docker-compose.yaml file is located, and run the command

- docker-compose up

The request data is received in JSON format:

        {  
            "title": "someTitle",
            "author": "author1, author2, author3, ...",
            "publishing": "somePublishing"
        }  

The application is configured to run with the Linux operating system
(see Dockerfile, GOOD and GOARCH variables). To run on another OS, you need to change them.