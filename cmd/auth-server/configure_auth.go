package main

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"

	"auth/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureAPI(api *operations.AuthAPI) {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = httpkit.JSONConsumer()

	api.JSONProducer = httpkit.JSONProducer()

}
