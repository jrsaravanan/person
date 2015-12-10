
# Auth API

Authentication and Authorization API 

REST Web Service provides interface to authentication and authentication. It is integrated with RIL Active Directory. Can validated given user is available and 'active' in the AD server or not. It will generate a token on successful login and token removed will be removed after 30 minutes  if there are no action.

 - Validate user in AD and create a UUID  
 - Maintain authentication session validate user for any application
 - Provides Roles and Permission for CMDB

## Limitations
 - No Generic Roles/Permission handling
 - No Power user / admin functionalities avaiable :  super user to manage all other users
 - List / In validate avaialable and it will NOT be exposed 

## Authentication Token Invalidation
Authentication token will be invalidated every 30 mins. If there are no activity or touch in the session for 20 mins it will be considered as inactive session.

## How to integrate Auth API in a application ?
Here are the steps
    - Login or authenticate using v1/auth/x 
    - Once you logged in you will be provided with x-auth-token
    - Use x-auth-token as a key 
    - validate the it is valid x-auth-token or not using v1/auth/x/{x-auth-token}

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

 - v1/auth/ping : ping : GET
 - v1/auth/x : login : POST
 - v1/auth/{x-auth-token}/roles : get roles : GET
 - v1/auth/x/{x-auth-token} : touch user is active :GET

 Admin Operations
 > It requies a key  to get your key , reach out development team for key
 - v1/list : list tokens : GET
 - /v1/auth/roles : Add new roles : POST

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


### Roles v1/auth/{x-auth-token}/roles
To get available CMDB roles for that specific user. 

```
Request 
http://localhost:9090/v1/auth/5f70b7b6-7f96-40cd-a317-222a1621efbf/roles
Http Method : POST

Response:

Success :
    Status Code: 200 OK
    {
        "Id": 91,
        "userName": "test",
        "team": "test",
        "Roles":
        [
            {
                "role_Name": "AUTOMATION",
                "permission":
                [
                    "access_site",
                    "access_pe",
                    "create_project",
                    "view_project",
                    "create_subnet",
                    "view_subnet",
                    "create_os_profile",
                    "view_os_profile"
                ]
            }
        ]
    }

Failure
    Status Code : 404 Not Found
    Invalid token xxxx-xxxx-xxx
```

### List v1/auth/list
Returns all the available active tokens.

```
Request 
http://localhost:9090/v1/auth/list
Http Method : GET

Response:

Sucess :
     Status Code: 200 OK
    {
        "a90b79a4-4baa-4f98-9c89-87038b60aa0c":
        {
            "userName": "saravanan",
            "x-auth-token": "a90b79a4-4baa-4f98-9c89-87038b60aa0c",
            "tokenCreateAt": "2015-12-09T17:46:09.6391637-05:00",
            "lastUpdatedTime": "2015-12-09T17:46:09.6391637-05:00"
        },
        "d5c3f742-9a5d-444d-aa60-e51a8f84d8ba":
        {
            "userName": "test",
            "x-auth-token": "d5c3f742-9a5d-444d-aa60-e51a8f84d8ba",
            "tokenCreateAt": "2015-12-09T17:46:19.9354148-05:00",
            "lastUpdatedTime": "2015-12-09T17:46:19.9354148-05:00"
        }
    }

Failure
     Status Code: 503 
```

### Touch v1/auth/x/{x-auth-token}
touch the token to keep it live

```
Request 
http://localhost:9090/v1/auth/x/79aa10c0-6538-4b2c-ac19-dc6e7f00a3fc
Http Method : GET

Response:

Sucess :
     Status Code: 200 OK
    "79aa10c0-6538-4b2c-ac19-dc6e7f00a3fc"

Failure
     Status Code: 404 Not Found
     Invalid token 79aa10c0-6538-4b2c-ac19-dc6e7f00a3fcsd
```


{"userName":"testxyz","team":"testxyx","domain":"local" ,"Roles":[{"Id" : 10}] }


### AddRole v1/auth/roles
add new roles to user. This method has some known issuse

```
Request 
http://localhost:9090/v1/auth/roles
{"userName":"test812q","team":"IDC Test Team","Roles":{"Id" :121210} }
Http Method : POST

Response:

Sucess :
    

    {
        "Id": 117,
        "userName": "test82",
        "team": "IDC Test Team",
        "Roles":
        {
            "roleName": "NOC SERVER"
        }
    }



Failure
     Status Code: 400 Bad Request
     
```


