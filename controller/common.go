//TODO: I believe the better name for this file is commonctrl.go

package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	. "auth/log"
	"auth/token"
	"auth/types"
)

type (

	//ICommonController common interface
	//list of common functions
	ICommonController interface {
		InitDB()
		Ping(w http.ResponseWriter, r *http.Request)
		Login(w http.ResponseWriter, r *http.Request)
		Roles(w http.ResponseWriter, r *http.Request)
		ListTokens(w http.ResponseWriter, r *http.Request)
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
		doErrorTranslation(w, err)
		return
	}

	u, err := h.xauth.LoginUser(*l)
	if err != nil {
		doErrorTranslation(w, err)
		return
	}

	if err = json.NewEncoder(w).Encode(u); err != nil {
		doErrorTranslation(w, err)
		return
	}
}

// Roles get roles for given x-auth-token
func (h *CommonController) Roles(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uuid := vars["x-auth-token"]
	Logger.Debugf("get roles for %s", uuid)

	ur, err := h.xauth.AuthRoles(uuid)
	if err != nil {
		doErrorTranslation(w, err)
		return
	}

	if err = json.NewEncoder(w).Encode(ur); err != nil {
		doErrorTranslation(w, err)
		return
	}
}

// ListTokens Authentication Tokens
func (h *CommonController) ListTokens(w http.ResponseWriter, r *http.Request) {
	list := h.xauth.ListTokens()
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//translate error codes on error
//
func doErrorTranslation(w http.ResponseWriter, err error) {

	Logger.Errorf("Message %s - error - %+v", err.Error(), err)
	if strings.EqualFold(err.Error(), "record not found") {
		http.Error(w, "Invalid user name or password", http.StatusNotFound)
	} else if strings.Contains(err.Error(), "Invalid token") {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// InvalidateToken
func InvalidateToken() {
	token.InvalidateTokens()
}

//InitDB initalize DB
func (h *CommonController) InitDB() {
	h.xauth.InitDB()
}
