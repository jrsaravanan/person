//TODO: I believe the better name for this file is commonctrl.go

package controller

import (
	. "auth/log"
	"net/http"
	"time"
)

type (

	//ICommonController common interface
	//list of common functions
	ICommonController interface {
		Ping(w http.ResponseWriter, r *http.Request)
	}

	//CommonController struct
	CommonController struct{}
)

//Ping to verify the service is up or not
//service testing method
func (h *CommonController) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("success"))
	//w.Header().Set("location", r.URL.Path)
	w.WriteHeader(http.StatusAccepted)

}

//WithLogging log the request and response
func WithLogging(w http.ResponseWriter, r *http.Request,
	executeFunc func(w http.ResponseWriter, r *http.Request)) {

	startTime := time.Now()
	executeFunc(w, r)
	endTime := time.Now()
	Logger.Debug("ElapsedTime in seconds:", endTime.Sub(startTime))

}
