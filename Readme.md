
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

### Test
```
go test -coverprofile cover.out 
go tool cover -html=cover.out -o cover.html
```

## To Build
```
On your woking directory 
git clone git@10.137.2.164/RJILIDCAutomation/auth.git
go install auth/...
```

### Swagger : Under progress 
```
swagger -apiPackage="Authentication API" -mainApiFile="auth/main.go"
```

## Endpoints


http://localhost:9090/v1/auth
{"userName":"test","password":"test","domain":"local"}

Success :


    {
        "userName": "test",
        "x-auth-token": "e8e5d0e3-8b3d-49ef-a1d0-b36d9cd10a9d",
        "tokenCreateAt": "2015-12-09T09:53:26.1163404-05:00",
        "lastUpdatedTime": "2015-12-09T09:53:26.1163404-05:00"
    }

