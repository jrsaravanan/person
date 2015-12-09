package token

import (
	. "auth/log"
	"auth/types"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CreateToken create new token
func TestCreateToken(t *testing.T) {

	uuid := CreateToken()
	fmt.Println("uuid", uuid)
	assert.NotEmpty(t, uuid, "UUID is empty")
}

// CreateToken create new token
func TestUsedToken(t *testing.T) {

	uuid := CreateToken()
	token[uuid] = types.LoginUser{UserName: "test_user"}
	token := findExistingToken("test_user")
	assert.Equal(t, uuid, token, "wrong token")

	token = findExistingToken("test")
	assert.Equal(t, "", token, "wrong token")
}

func TestAuthentication(t *testing.T) {

	err := NewDataAccess()
	defer dbConn.Close()
	assert.NoError(t, err, "db connection failed")

	auth := AuthToken{}
	l, err := auth.LoginUser(types.LoginUser{UserName: "test", Password: "test", Domain: "local"})
	assert.NoError(t, err, "expected no record error")

	Logger.Debugf("user values %+v", l)
	assert.Equal(t, "test", l.UserName, "user name should be same ")
	assert.NotEmpty(t, l.AuthToken, "auth token should be generated")
	previousToken := l.AuthToken

	// test alread logged user
	l, err = auth.LoginUser(types.LoginUser{UserName: "test", Password: "test", Domain: "local"})
	assert.NoError(t, err, "reqired no record expcetion")
	assert.Equal(t, previousToken, l.AuthToken, "it should not create token for second time")

}

func TestRoles(t *testing.T) {

	err := NewDataAccess()
	defer dbConn.Close()
	assert.NoError(t, err, "db connection failed")

	auth := AuthToken{}
	l, err := auth.LoginUser(types.LoginUser{UserName: "test", Password: "test", Domain: "local"})
	assert.NoError(t, err, "expected no record error")

	Logger.Debugf("user values %+v", l)
	assert.Equal(t, "test", l.UserName, "user name should be same ")
	assert.NotEmpty(t, l.AuthToken, "auth token should be generated")

	r, err := auth.AuthRoles(l.AuthToken)
	assert.NoError(t, err, "should not be any error while getting the roles")
	Logger.Debugf("Roles %+v", r)
}
