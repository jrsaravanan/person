
# Auth API

Authentication and Authorization API 

REST Web Service provides interface to authentication and authentication. It is integrated with RIL Active Directory. Can validated given user is available and 'active' in the AD server or not. It will generate a token on successful login and token removed will be removed after 30 minutes  if there are no action.

 1 Validate user in AD and create a UUID  
 2 Maintain authentication session validate user for any application
 3 Provides Roles and Permission for CMDB

## Limitations
 1 No Generic Roles/Permission handling
 2 No Power user / admin functionalities avaiable :  super user to manage all other users
 3 List / In validate avaialable and it will NOT be exposed 


## To Build
```
On your working directory 
git clone git@10.137.2.164/RJILIDCAutomation/auth.git
go install auth/...
```

### Test
```
go test -coverprofile cover.out 
go tool cover -html=cover.out -o cover.html
```


### Swagger : Under progress 
```
swagger -apiPackage="Authentication API" -mainApiFile="auth/main.go"
```

## API Endpoint and Operations
List of supported resources, use with base url 

 - v1/auth/ping
 - v1/auth/x

### Auth v1/auth/ping
To check the service is up and running.

```
Request
http://localhost:9090/v1/auth/ping
Http Method : GET

Success :
	Status Code: 200 OK
Failure :
    No Response : Service Not Available
    Status Code : 50x Service Not Available / Internal Error
```

### Auth v1/auth
Validate the login informations and provides authentication token on sucessfull login.
It supports RIL AD (domain : in) and database authentication. By default the user name and passowrd will be verified in Active Directory.If you want to enable  database authentication you should send 'domain' flag in your request. 
> domain : local - enable database authentication
i.e : {"userName":"test","password":"test","domain":"local"}

```
Request 
http://localhost:9090/v1/auth/x
{"userName":"test","password":"test"}
ttp Method : POST

Response:

Success :
	Status Code: 200 OK
    {
        "userName": "test",
        "x-auth-token": "e8e5d0e3-8b3d-49ef-a1d0-b36d9cd10a9d",
        "tokenCreateAt": "2015-12-09T09:53:26.1163404-05:00",
        "lastUpdatedTime": "2015-12-09T09:53:26.1163404-05:00"
    }

Failure :
    Status Code: 404 Not Found
    Invalid user name or password
```    


