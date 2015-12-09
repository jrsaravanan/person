//TODO: I believe the better name for this file is commonctrl.go

package controller

import (
	"encoding/json"
	"net/http"
	"time"

	. "auth/log"
	"auth/token"
	"auth/types"
)

type (

	//ICommonController common interface
	//list of common functions
	ICommonController interface {
		Ping(w http.ResponseWriter, r *http.Request)
		Login(w http.ResponseWriter, r *http.Request)
		InitDB()
	}

	//CommonController struct
	CommonController struct {
		xauth token.AuthToken
	}
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

// Login authenication interface
func (h *CommonController) Login(w http.ResponseWriter, r *http.Request) {

	var l *types.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.xauth.LoginUser(*l)
	if err != nil {
		Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if er := json.NewEncoder(w).Encode(u); er != nil {
		Logger.Error(err.Error())
		http.Error(w, er.Error(), http.StatusInternalServerError)
	}
}

//InitDB initalize DB
func (h *CommonController) InitDB() {
	token.NewDataAccess()
}
