//Package user provides list of function to authentication

//Search based on https://github.com/nmcclain/ldap/blob/master/examples/search.go
//Copyright 2014 The Go Authors. All rights reserved.
//Use of this source code is governed by a BSD-style
//license that can be found in the LICENSE file.

package token

import (
	. "auth/log"
	"errors"
	"fmt"

	"github.com/nmcclain/ldap"
)

var (
	ldapServer        = "in.ril.com"
	ldapPort   uint16 = 389
	baseDN            = "dc=in,dc=ril,dc=com"
	attributes        = []string{"memberof"}
)

const (
	//NoServerConnection used for connection error
	NoServerConnection string = "Unable to connect to LDAP server"

	//InvalidCredentials used for credential error
	InvalidCredentials string = "Invalid Credentials"

	//NoSearchData used for not able to find the details after binding
	NoSearchData string = "Check user is active"
)

//Authenticate check the given user has permission to bind and search same user
//it will return true when succesfully logged in
func Authenticate(user string, passwd string) (bool, error) {

	Logger.Debug("login in ", user)
	filter := "(&(objectClass=user)(sAMAccountName=" + user + "))"
	user = "in\\" + user
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
	if err != nil {
		Logger.Errorf("ERROR: %s\n", err.Error())
		return false, errors.New(NoServerConnection)
	}
	defer l.Close()
	// l.Debug = true

	err = l.Bind(user, passwd)
	if err != nil {
		Logger.Errorf("ERROR: Cannot bind: %s\n", err.Error())
		return false, errors.New(InvalidCredentials)
	}
	search := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		attributes,
		nil)

	sr, err := l.Search(search)
	if err != nil {
		Logger.Errorf("ERROR: %s\n", err.Error())
		return false, errors.New(NoSearchData)
	}

	Logger.Debugf("Search: %s -> num of entries = %d\n", search.Filter, len(sr.Entries))
	login := len(sr.Entries) > 0

	return login, nil
}
