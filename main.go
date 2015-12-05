/**
IRIS Service main package

*/
package main

import (
	"flag"
	"net/http"
	"strings"
	"time"

	"auth/log"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vharitonsky/iniflags"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	isJsonRequest := false
	w.Header().Set("Content-Type", "text/html")

	if acceptHeaders, ok := r.Header["Accept"]; ok {
		for _, acceptHeader := range acceptHeaders {
			if strings.Contains(acceptHeader, "json") {
				isJsonRequest = true
				break
			}
		}
	}

	if isJsonRequest {
		w.Write([]byte(resourceListingJson))
	} else {
		http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
	}

}

func ApiDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	apiKey := strings.Trim(r.RequestURI, "/")

	if json, ok := apiDescriptionsJson[apiKey]; ok {
		w.Write([]byte(json))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

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
		httpPort      string
		globalMux     = http.NewServeMux()
		staticContent = flag.String("staticPath", "./swagger-ui", "Path to folder with Swagger UI")
	)

	flag.StringVar(&httpPort, "http.port", "9090", "http listening port")

	//load application properties from *.ini file
	//Init data access
	iniflags.Parse()

	//all route operations added here
	httpRouter := mux.NewRouter()
	httpRouter.HandleFunc("/", IndexHandler).Methods("GET")

	//api := operations.AuthAPI{}

	//globalMux.Handle("/", api.Serve())
	globalMux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir(*staticContent))))
	//globalMux.Handle("/swagger-ui/", http.FileServer(http.Dir("/swagger-ui")))
	globalMux.Handle("/v1/auth/", buildRouter())
	log.Logger.Info("Authentication Server listening on port - ", httpPort)

	//to handle CORS support
	//Angular JS may use OPTION for PUT request , it is handled by CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	corsHandler := c.Handler(globalMux)
	//responseHandler := preJSONProcessor(corsHandler)

	//http.HandleFunc("/", IndexHandler)
	//http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir(*staticContent))))

	for apiKey, _ := range apiDescriptionsJson {
		http.HandleFunc("/"+apiKey+"/", ApiDescriptionHandler)
	}

	//start server
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        context.ClearHandler(corsHandler),
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Logger.Fatal(s.ListenAndServe())

}
