# Simple Rest Api 

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/2cde77c93e034ff6b068bc9ec198f4ea)](https://app.codacy.com/manual/AdrianMendez1199/simple-crud-go-sql?utm_source=github.com&utm_medium=referral&utm_content=AdrianMendez1199/simple-crud-go-sql&utm_campaign=Badge_Grade_Settings)

## Requirements
 - GO `go version` go version go1.14.3 darwin/amd64 or later
 
 - postgresql `postgres --version` postgres (PostgreSQL) 11.5 or later

 ## Run
 - go build 
 - ./simple-crud-go-sql 
 

 ## Endpoints
 - /student POST ***Request:***
 ```json
 {
    "name": "username",
    "age": 22,
    "active": true
 }
 ```
***Response:***
```json
{
    "status": "OK",
    "message": "user created"
}
```

----
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