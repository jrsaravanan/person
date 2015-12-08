package token

import (
	. "auth/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUserPermissions(t *testing.T) {

	Logger.Debug("test permission")
	err := NewDataAccess()
	defer dbConn.Close()
	assert.NoError(t, err, "db connection failed")

	u := UserRepository{}
	list, err := u.FindUserPermissions("Saravanan.Renganatha")
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

func TestRoles(t *testing.T) {

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

	repo.UpdateRoles("test", "test", 3)

}
