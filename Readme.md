
# Auth API

Authentication and Authorization API 

REST Web Service provides interface to authentication and authentication. It is integrated with RIL Active Directory. Can validated given user is available and 'active' in the AD server or not. It will generate a token on successful login and token removed will be removed after 30 minutes  if there are no action.

 1 Validate user in AD and create a UUID  
 2 Maintain authentication session validate user for any application
 3 Provides Roles and Permission for CMDB

## Limitations
 1 No Generic Roles/Permission handling
 2 No Admin Interface (User CRUD operations not supported) for CMDB Roles

### Test
```
go test -coverprofile cover.out 
go tool cover -html=cover.out -o cover.html
...

### Swagger : Under progress 
```
swagger -apiPackage="Authentication API" -mainApiFile="auth/main.go"
...