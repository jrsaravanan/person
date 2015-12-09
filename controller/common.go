//TODO: I believe the better name for this file is commonctrl.go

package controller

import (
	. "auth/log"
	"auth/types"
	"encoding/json"
	"net/http"
	"time"
)

type (

	//ICommonController common interface
	//list of common functions
	ICommonController interface {
		Ping(w http.ResponseWriter, r *http.Request)
		Login(w http.ResponseWriter, r *http.Request)
	}

	//CommonController struct
	CommonController struct{}
)

//Ping to verify the service is up or not
//service testing method
func (h *CommonController) Ping(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("location", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

//WithLogging log the request and response
func WithLogging(w http.ResponseWriter, r *http.Request,
	executeFunc func(w http.ResponseWriter, r *http.Request)) {

	startTime := time.Now()
	executeFunc(w, r)
	endTime := time.Now()
	Logger.Debug("ElapsedTime in seconds:", endTime.Sub(startTime))

}

// Login authenrication interface
func (h *CommonController) Login(w http.ResponseWriter, r *http.Request) {

	var l *types.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
