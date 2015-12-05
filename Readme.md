go test -coverprofile cover.out and then
go tool cover -html=cover.out -o cover.html


swagger -apiPackage="Authentication API" -mainApiFile="auth/main.go"