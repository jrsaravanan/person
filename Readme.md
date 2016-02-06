
# Person API

Sample API person API

Person sample REST API 


## To Build
```
go install person/...
```

### Test
```
go test -coverprofile cover.out 
go tool cover -html=cover.out -o cover.html
```

### Run Local
```
<your workspace>\src\person>go run person.go -config=config/person-config-local.ini

```


### Swagger : Under progress 
```
swagger -apiPackage="Authentication API" -mainApiFile="person/main.go"
```

## API Endpoint
List of supported resources, use with base url 

 - v1/person/ping : ping : GET
 

### Person v1/person/ping
To check the service is up and running.

```
Request
http://localhost:9090/v1/person/ping
Http Method : GET

Success :
	Status Code: 200 OK
Failure :
    No Response : Service Not Available
    Status Code : 50x Service Not Available / Internal Error
```
