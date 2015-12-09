package token

import (
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
		InitDB()
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

	}

	return
}

// Authenticate login interface method
func (a *AuthToken) Authenticate(l types.LoginUser) {

}

func findExistingToken(userName string) string {

	for key, u := range token {
		if u.UserName == userName {
			return key
		}
	}
	return ""
}

func (a *AuthToken) InitDB() {
	a.InitDB()
}
