package token

import (
	"bytes"
	"flag"
	"fmt"
	"strconv"
	"strings"

	. "auth/log"
	"auth/types"

	// load sql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	dbConn gorm.DB
	err    error
)

var (
	//CMDB configurations
	dbschema string
	dbip     string
	dbport   string
	user     string
	secret   string
)

func init() {

	flag.StringVar(&dbschema, "db.schema", "cmdb", "db schema name")
	flag.StringVar(&dbip, "db.ip", "127.0.0.1", "database ip address")
	flag.StringVar(&dbport, "db.port", "3306", "database ip address")
	flag.StringVar(&user, "db.user", "test", "database user name")
	flag.StringVar(&secret, "db.password", "test123", "database password")
}

type (

	// IUserRepository interface
	// set of  CURD operations on user and roles
	IUserRepository interface {
		FindUserPermissions(user string) ([]string, error)
		LoginUser(user string, password string) (b bool, err error)
		UpdateRoles(user string, team string, roleId string)
	}

	//UserRepository empty struct
	UserRepository struct{}
)

// NewDataAccess create dbconnection
func NewDataAccess() (err error) {

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprint(user, ":", secret, "@tcp(", dbip, ":", dbport, ")/", dbschema))
	Logger.Debug("\n MySQL Database Connection String :", buffer.String())
	dbURL := buffer.String()
	dbConn, err = gorm.Open("mysql", dbURL)
	if err != nil {
		return
	}

	dbConn.DB().Ping()
	dbConn.DB().SetMaxIdleConns(10)
	dbConn.DB().SetMaxOpenConns(20)
	dbConn.SingularTable(true)
	dbConn.LogMode(true)
	return

}

// LoginUser get login user
func (u *UserRepository) LoginUser(user string, password string) (b bool, err error) {

	usr := new(types.LoginUser)
	err = dbConn.Find(&usr, types.LoginUser{UserName: user, Password: password}).Error
	if err != nil {
		Logger.Error(err.Error())
		return
	}

	if usr != nil {
		b = true
	}
	Logger.Debugf("user %+v", user)
	return
}

// Roles get login user
// TODO : Object releation mapping is not working
// may be some sily mistake , not able to figure it out
// going for dirty fix - should be fixed after first release
/*func (u *UserRepository) Roles(user string) (usr *types.User, err error) {

	usr = &types.User{UserName: user}
	var roles []types.Role

	err = dbConn.Find(&usr, types.User{UserName: user}).Error
	err = dbConn.Model(&usr).Related(&roles).Error
	if err != nil {
		Logger.Error(err.Error())
		return
	}
	Logger.Debugf("User details %+v", usr)
	Logger.Debugf("Roles %+v", roles)
	return
}*/

// Roles get login user
func (u *UserRepository) Roles(user string) (usr *types.User, err error) {

	usr = new(types.User)
	err = dbConn.Find(&usr, types.User{UserName: user}).Error
	if err != nil {
		Logger.Error("user details failed . ", err.Error())
		return
	}

	query := " SELECT role_name  FROM t_user_role ur ,  t_role r WHERE ur.role_id = r.role_id  " +
		"AND  ur.user_id = " + strconv.Itoa(usr.ID)
	Logger.Debug("query - ", query)

	rows, err := dbConn.DB().Query(query)
	defer rows.Close()
	if err != nil {
		Logger.Error("Error ", err.Error())
		return
	}
	var roleName string
	rows.Next()
	rows.Scan(&roleName)

	Logger.Debugf("User - %v Role Name -  %v ", usr.UserName, roleName)
	usr.Roles[0] = types.Role{RoleName: roleName}

	return
}

//FindUserPermissions to get permission list for given user
func (u *UserRepository) FindUserPermissions(user string) (permissions []string, err error) {

	query := "SELECT  p.perm_key FROM  t_role_permission AS rp LEFT JOIN " +
		" t_user_role AS ur ON rp.role_id = ur.role_id " +
		" LEFT JOIN t_permission AS p ON rp.perm_id =p.perm_id " +
		" LEFT JOIN login_user AS u ON ur.user_id = u.user_id" +
		" WHERE  u.username = '" + user + "' "

	Logger.Debug("query : ", query)

	rows, err := dbConn.DB().Query(query)
	if err != nil {
		Logger.Error("Error ", err.Error())
		return
	}

	var column string
	for rows.Next() {
		err = rows.Scan(&column)
		if err != nil {
			Logger.Error("Error ", err.Error())
			return
		}
		permissions = append(permissions, column)
	}

	return
}

func (u *UserRepository) UpdateRoles(user string, team string, roleId int) {

	//find user availablity
	var usr types.User
	tx := dbConn.Begin()
	err = tx.Find(&usr, types.User{UserName: user}).Error
	if err != nil && !strings.EqualFold(err.Error(), "record not found") {
		tx.Rollback()
		Logger.Error(err.Error())
		return
	}

	Logger.Debugf("user value - %+v", usr)

	//if user not available , create new user
	if usr.UserName == "" {
		err = tx.Save(&types.User{UserName: user, Team: team}).Error
		if err != nil {
			tx.Rollback()
			Logger.Error("Saving user failed ", err.Error())
		}

	}

	tx.Commit()
	tx.Close()
	//tx.Commit()
	Logger.Debugf("New user %s created ", user)
}
