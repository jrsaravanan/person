package token

import (
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
