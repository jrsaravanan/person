package token

import (
	"flag"
	"fmt"
	"sync"

	. "auth/log"
	"auth/types"
	"time"

	"github.com/pborman/uuid"
)

var (
	lock  = &sync.Mutex{}
	token = make(map[string]types.LoginUser)
	sto   float64
	//umap  = make(map[string]types.User)
)

type (

	//IAuthToken token interface
	IAuthToken interface {
		LoginUser(l types.LoginUser) (usr types.LoginUser, err error)
		AuthRoles(xauth string) (usr *types.User, err error)
		InitDB() (err error)
		TouchToken(xauth string) (err error)
		UpdateRoles(u types.User) (usr *types.User, err error)
	}

	//AuthToken token
	AuthToken struct {
		repo UserRepository
	}
)

func init() {
	flag.Float64Var(&sto, "session.timeout", 2.0, "auth token session time out ")
}

// CreateToken create new token
func CreateToken() string {
	lock.Lock()
	defer lock.Unlock()
	id := uuid.New()
	Logger.Debugf("new uuid created %s", id)
	return id
}

// LoginUser login user
func (a *AuthToken) LoginUser(l types.LoginUser) (usr types.LoginUser, err error) {

	var login bool
	if l.Domain == "local" {
		login, err = a.repo.LoginUser(l.UserName, l.Password)
	} else {
		login, err = Authenticate(l.UserName, l.Password)
	}
	Logger.Debug("login status ", login)

	if err != nil {
		Logger.Error("login failed ", err.Error())
		return
	}

	//if user already logged in send the same x-auth-token
	//update the LastUpdatedTime to now
	authToken := findExistingToken(l.UserName)
	usr = token[authToken]
	if usr == (types.LoginUser{}) {
		authToken = CreateToken()
		usr = types.LoginUser{
			UserName:        l.UserName,
			AuthToken:       authToken,
			TokenCreateAt:   time.Now(),
			LastUpdatedTime: time.Now(),
		}
		token[authToken] = usr

	} else {
		usr.LastUpdatedTime = time.Now()
	}

	return
}

// Authenticate login interface method
func (a *AuthToken) Authenticate(l types.LoginUser) {

}

// AuthRoles get roles for the given xauth token
// return invalid token error for unavailable user token
func (a *AuthToken) AuthRoles(xauth string) (r *types.User, err error) {

	u := token[xauth]
	if u == (types.LoginUser{}) {
		err = fmt.Errorf("Invalid token %s", xauth)
		return
	}

	r, err = a.repo.Roles(u.UserName)
	if err != nil {
		return
	}
	return
}

func findExistingToken(userName string) string {

	for key, u := range token {
		if u.UserName == userName {
			return key
		}
	}
	return ""
}

//ListTokens list avaiable tokens
func (a *AuthToken) ListTokens() map[string]types.LoginUser {
	return token
}

// TocuhToken update active timestamp
func (a *AuthToken) TocuhToken(xauthToken string) (err error) {
	u := token[xauthToken]

	if u == (types.LoginUser{}) {
		err = fmt.Errorf("Invalid token %s", xauthToken)
		return
	}
	Logger.Debugf("updated user %+v", u)
	// touch the time
	u.LastUpdatedTime = time.Now()
	return
}

// InvalidateTokens find auth session are valid
func InvalidateTokens() {
	lock.Lock()
	defer lock.Unlock()

	for key, u := range token {
		if findTimeOut(u) {
			delete(token, key)
		}
	}
}

// find if the user is active or not
func findTimeOut(usr types.LoginUser) bool {
	now := time.Now()
	duration := now.Sub(usr.LastUpdatedTime)
	if duration.Minutes() > sto {
		return true
	}
	return false
}

// InitDB initalize DB
// wrapper class
func (a *AuthToken) InitDB() (err error) {
	err = NewDataAccess()
	return
}

// UpdateRoles get roles for the given xauth token
// return invalid token error for unavailable user token
func (a *AuthToken) UpdateRoles(u types.User) (r *types.User, err error) {

	Logger.Debugf("user %s , team %s , roles %d", u.UserName, u.Team, u.Roles.ID)
	a.repo.UpdateRoles(u.UserName, u.Team, u.Roles.ID)
	r, err = a.repo.Roles(u.UserName)
	if err != nil {
		return nil, err
	}
	return r, nil
}
