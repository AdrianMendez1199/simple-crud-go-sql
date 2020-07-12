# Simple Rest Api 

## Requirements
 - GO `go version` go version go1.14.3 darwin/amd64 or later
 
 - postgresql `postgres --version` postgres (PostgreSQL) 11.5 or later

 ## Run
 - go build 
 - ./simple-crud-go-sql 
 

 ## Endpoints
 - /student POST 
 ```json
 {
    "name": "username",
    "age": 22,
    "active": true
 }
 ```

 - /students GET
  ***response:***
  ---
 ```json
 [
    {
        "id": 1,
        "name": "username",
        "age": 30,
        "active": true,
        "createdAt": "0001-01-01T00:00:00Z",
        "updatedAt": "0001-01-01T00:00:00Z"
    }
 ]
 ```

- /student/{id} GET ***response:***
***
```json
{
    "id": 1,
    "name": "name",
    "age": 30,
    "active": true,
    "createdAt": "2020-07-10T09:46:14.247305Z",
    "updatedAt": "0001-01-01T00:00:00Z"
}
```