package token

import (
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
	//umap  = make(map[string]types.User)
)

type (

	//IAuthToken token interface
	IAuthToken interface {
		LoginUser(l types.LoginUser) (usr types.LoginUser, err error)
		AuthRoles(xauth string) (usr *types.User, err error)
		InitDB() (err error)
	}

	//AuthToken token
	AuthToken struct {
		repo UserRepository
	}
)

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

// InitDB initalize DB
// wrapper class
func (a *AuthToken) InitDB() (err error) {
	err = NewDataAccess()
	return
}
