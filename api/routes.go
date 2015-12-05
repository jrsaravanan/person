package api

import (
	"auth/controller"
	"net/http"

	"github.com/gorilla/mux"
)

//BuildRouter building all routes
func BuildRouter() *mux.Router {
	//init routers and controllers
	apiRouter := mux.NewRouter()
	commonController := new(controller.CommonController)

	//Add Routes and controllers
	AddPingRoute(apiRouter, commonController)

	return apiRouter
}

// AddPingRoute to check service avaialblity
// @Title Ping
// @Description ping auth service
// @Accept  json
// @Success 200 string string
// @Failure 404
// @Router /v1/auth/ping [get]
func AddPingRoute(apiRouter *mux.Router, h controller.ICommonController) {
	apiRouter.HandleFunc("/v1/auth/ping", func(rw http.ResponseWriter, req *http.Request) {
		controller.WithLogging(rw, req, h.Ping)
	}).Methods("GET")
}
