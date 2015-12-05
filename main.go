/**
IRIS Service main package

*/
package main

import (
	"flag"
	"net/http"
	"time"

	"auth/api"
	. "auth/log"

	"github.com/gorilla/context"
	"github.com/rs/cors"
	"github.com/vharitonsky/iniflags"
)

//postJSONProcessor add http header for all response
func preJSONProcessor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func main() {

	var (
		httpPort  string
		globalMux = http.NewServeMux()
	)

	flag.StringVar(&httpPort, "http.port", "9090", "http listening port")

	//load application properties from *.ini file
	//Init data access
	iniflags.Parse()

	globalMux.Handle("/v1/auth/", api.BuildRouter())
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
	responseHandler := preJSONProcessor(corsHandler)

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
