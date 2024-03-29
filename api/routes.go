//Package api list of APIS
// @Contact saravanan.renganathan@gmail.com
// @BasePath http://127.0.0.1
package api

import (
	"net/http"
	"person/controller"

	"github.com/gorilla/mux"
)

//BuildAuthRouter building all routes
func BuildAuthRouter() *mux.Router {

	//init routers and controllers
	apiRouter := mux.NewRouter()
	commonController := new(controller.CommonController)

	//Add Routes and controllers
	AddPingRoute(apiRouter, commonController)

	return apiRouter
}

 

// PreJSONProcessor add http header for all response
// intercept every request
func PreJSONProcessor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// AddPingRoute to check service avaialblity
// @Title Ping
// @Description ping auth service
// @Accept  json
// @Success 200 string string
// @Failure 404 string string
// @Router /v1/auth/ping [get]
func AddPingRoute(apiRouter *mux.Router, h controller.ICommonController) {
	apiRouter.HandleFunc("/v1/person/ping", func(rw http.ResponseWriter, req *http.Request) {
		controller.WithLogging(rw, req, h.Ping)
	}).Methods("GET")
}
