
package main
//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "",
    "swaggerVersion": "1.2",
    "apis": [
        {
            "path": "/v1",
            "description": "ping auth service"
        }
    ],
    "info": {}
}`
var apiDescriptionsJson = map[string]string{"v1":`{
    "apiVersion": "",
    "swaggerVersion": "1.2",
    "basePath": "{{.}}",
    "resourcePath": "/v1",
    "apis": [
        {
            "path": "/v1/auth/ping",
            "description": "ping auth service",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Ping",
                    "type": "string",
                    "items": {},
                    "summary": "ping auth service",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "string",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        }
    ]
}`,}
