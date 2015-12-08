package token

import (
	"sync"

	. "auth/log"
	"auth/types"

	"github.com/pborman/uuid"
)

var (
	lock = &sync.Mutex{}
	m    = make(map[string]types.User)
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
func LoginUser(user string, password string) (u types.User, err error) {
	return
}
