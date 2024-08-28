
## Employee management CRUD Application 

### **Project Overview**:
This is a Golang based CRUD application which leverages the gorilla-mux framework and MongoDB database to create RESTful API for the Employee management system. 

### **Workflow**:
The project workflow includes: 
* Extracting the data from an endpoint.
* Storing the data in MongoDB.
* Performing CRUD operations on the data.


### **Functionalities**:
The project functionalities are: 
* Fetching data from an endpoint.

* Inserting data into MongoDB.
* Reading all data from MongoDB.
* Reading a single data entry from MongoDB using an Object ID. Updating data in MongoDB.
* Deleting data from MongoDB.
* Retrieving data from MongoDB.

### **Pre-requisites** 
* Python (version 3.10.12 or later)
* Golang (version 1.22.5 or later)
* Docker (optional, but recommended for easy deployment): https://www.docker.com
* Make (version 3.81 or later) : https://formulae.brew.sh/formula/make

### **How to run locally**: 
Step 1: Clone the repository: 
```
git clone https://github.com/ch374n/log-analyser
```
Step 2: Install dependencies: 
```
go mod download
```
Step 3: Create a .env file in the project root directory (copy from .env.example if provided) and define environment variables:
```
APP_ENV=dev # Environment type
APP_PORT=8080 # Port on which the server will listen 
APP_MONGO_URI=mongsrv+ # Mongodb connection URL 
APP_MONGO_PORT=27017 # Mongodb server port

```
### **Http Endpoints**: 
| Endpoint       | HTTP Method | Description    |
|----------------|-------------|----------------|
| /employee      | GET         | GetAllHandler  |
| /employee      | POST        | CreateHandler  |
| /employee/{id} | GET         | GetHandlerById |
| /employee/{id} | PUT         | UpdateHandler  |
| /employee/{id} | DELETE      | DeleteHandler  |

### **API Usage with Swagger**:
Swagger documentation is available at http://localhost:8080/api-docs/.

Refer to the generated API definitions for details on endpoints and request/response
formats.

### **Building and running Docker image**:
Step 1: Build the docker image
```
docker build -t Log-Analysis-with-GenAI .
```
Step 2: Run the container: 
```
docker run -p $PORT:$PORT Log-Analysis-with-GenAI
```
