package token

import (
	"math/rand"
	"strconv"
	"testing"

	. "auth/log"
	"auth/types"

	"github.com/stretchr/testify/assert"
)

func TestFindUserPermissions(t *testing.T) {

	Logger.Debug("test permission")
	err := NewDataAccess()
	defer dbConn.Close()
	assert.NoError(t, err, "db connection failed")

	list, err := findUserPermissions("Saravanan.Renganatha")
	assert.NoError(t, err, "Getting permision failed")
	Logger.Debug("list - ", list)

}

func TestLoginUser(t *testing.T) {

	err := NewDataAccess()
	defer dbConn.Close()

	repo := UserRepository{}
	assert.NoError(t, err, "db connection failed")

	u, err := repo.LoginUser("saravanan", "saravanan")
	assert.NoError(t, err, "user not available")
	Logger.Debug("login flag ", u)
	assert.True(t, u, "user is available")

	u, err = repo.LoginUser("saravanan", "rete")
	assert.Error(t, err, "User not available")
	Logger.Debug("login flag ", u)
	dbConn.Close()

}

func TestUserRoles(t *testing.T) {

	err := NewDataAccess()
	defer dbConn.Close()
	assert.NoError(t, err, "db connection failed")

	repo := UserRepository{}

	r, err := repo.Roles("Saravanan.Renganatha")
	assert.NoError(t, err, "operation failed")
	Logger.Debugf("login flag %+v", r)
	assert.NotEmpty(t, r, "Roles Empty")
}

func TestUpdate(t *testing.T) {

	err := NewDataAccess()
	defer dbConn.Close()
	assert.NoError(t, err, "db connection failed")

	repo := UserRepository{}

	usr, err := repo.UpdateRoles("test", "test team", 3)
	assert.NoError(t, err, "failed update roles")
	assert.NotEmpty(t, usr, "user is not found")
}

func TestUpdateNewUser(t *testing.T) {

	err := NewDataAccess()
	defer dbConn.Close()
	assert.NoError(t, err, "db connection failed")

	repo := UserRepository{}
	num := rand.Intn(100)

	usr, err := repo.UpdateRoles("test"+strconv.Itoa(num), "test team", 3)
	assert.NoError(t, err, "failed update roles")
	assert.NotEmpty(t, usr, "user is not found")
}

func TestNoUser(t *testing.T) {

	err := NewDataAccess()
	defer dbConn.Close()
	assert.NoError(t, err, "db connection failed")

	auth := AuthToken{}
	_, err = auth.LoginUser(types.LoginUser{UserName: "testtest", Password: "testtest", Domain: "local"})
	assert.Error(t, err, "reqired no record expcetion")

}
