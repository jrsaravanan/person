package token

import (
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

// Login User
func TestExitingToken(t *testing.T) {

}
