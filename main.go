package main

import (
	"flag"
	"net/http"
	"time"

	"person/api"
	. "person/log"

	"github.com/gorilla/context"
	"github.com/rs/cors"
	"github.com/vharitonsky/iniflags"
)

func main() {

	var (
		httpPort  string
		globalMux = http.NewServeMux()
	)

	flag.StringVar(&httpPort, "http.port", "9090", "http listening port")

	//load application properties from *.ini file
	//Init data access
	iniflags.Parse()

	globalMux.Handle("/", api.BuildAPIRouter())
	globalMux.Handle("/v1/person/", api.BuildAuthRouter())
	Logger.Info("Authentication Server listening on port - ", httpPort)

	//to handle CORS support
	//Angular JS may use OPTION for PUT request , it is handled by CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	corsHandler := c.Handler(globalMux)
	responseHandler := api.PreJSONProcessor(corsHandler)

	//start server
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        context.ClearHandler(responseHandler),
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	Logger.Fatal(s.ListenAndServe())

}
