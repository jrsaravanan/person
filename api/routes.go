//Package api list of APIS
// @APIVersion 1.0.0
// @APITitle Auth API
// @APIDescription Authentication and autherization API
// @Contact saravanan.renganathan@ril.com
// @BasePath http://127.0.0.1
package api

import (
	"auth/controller"
	"flag"
	"net/http"

	"github.com/claudiu/gocron"
	"github.com/gorilla/mux"
)

var delay uint64

func init() {
	flag.Uint64Var(&delay, "session.scheduler", 10, "auth token session scheduler start time ")
}

//BuildAuthRouter building all routes
func BuildAuthRouter() *mux.Router {

	//init routers and controllers
	apiRouter := mux.NewRouter()
	commonController := new(controller.CommonController)
	commonController.InitDB()

	//Add Routes and controllers
	AddPingRoute(apiRouter, commonController)
	AddAuthRoute(apiRouter, commonController)
	AddRolesRoute(apiRouter, commonController)
	AddListTokenRoute(apiRouter, commonController)

	return apiRouter
}

//BuildAPIRouter building all routes
func BuildAPIRouter() *mux.Router {

	//init routers and controllers
	apiRouter := mux.NewRouter()
	apiRouter.Headers("Content-Type", "text/html")
	apiRouter.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/"))))
	return apiRouter
}

// StartScheduler start the scheduler
func StartScheduler() {

	gocron.Every(delay).Minutes().Do(controller.InvalidateToken)
	gocron.Start()
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
	apiRouter.HandleFunc("/v1/auth/ping", func(rw http.ResponseWriter, req *http.Request) {
		controller.WithLogging(rw, req, h.Ping)
	}).Methods("GET")
}

// AddAuthRoute login
// @Title Auth
// @Description ping auth service
// @Accept  json
// @Success 200 string string
// @Failure 404 string string
// @Router /v1/auth/x [post]
func AddAuthRoute(apiRouter *mux.Router, h controller.ICommonController) {
	apiRouter.HandleFunc("/v1/auth/x", h.Login).Methods("POST")
	apiRouter.HandleFunc("/v1/auth/x/{x-auth-token}", h.TouchToken).Methods("GET")
}

// AddRolesRoute get roles for given x-auth-token
// @Title GetRoles
// @Description returns roles and permission associated with x-auth-token
// @Accept  json
// @Success 200 string string
// @Failure 404 string string
// @Failure 503 string string
// @Router /v1/auth/{x-auth-token}/roles [post]
func AddRolesRoute(apiRouter *mux.Router, h controller.ICommonController) {
	apiRouter.HandleFunc("/v1/auth/{x-auth-token}/roles", h.Roles).Methods("GET")
}

// AddListTokenRoute get list of x-auth-token
// @Title GetTokensList
// @Description return authentication token list
// @Accept  json
// @Success 200 string string
// @Failure 503 string string
// @Router /v1/auth/{x-auth-token}/roles [post]
func AddListTokenRoute(apiRouter *mux.Router, h controller.ICommonController) {
	apiRouter.HandleFunc("/v1/auth/list", h.ListTokens).Methods("GET")
}
