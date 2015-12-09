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
		LoginUser(user string, password string) (b bool, err error)
		UpdateRoles(user string, team string, roleID int) (usr types.User, err error)
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
	Logger.Debugf("user %+v", usr)
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
	r := types.Role{RoleName: roleName}
	p, err := findUserPermissions(usr.UserName)
	if err != nil {
		Logger.Error("user roles failed ", err.Error())
		return
	}
	// usr.Roles[0].Permission = FindUserPermissions(usr.UserName)
	r.Permission = p
	usr.Roles[0] = r

	return
}

//findUserPermissions to get permission list for given user
func findUserPermissions(user string) (permissions []string, err error) {

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

// UpdateRoles add role to new user or exsiting user
func (u *UserRepository) UpdateRoles(user string, team string, roleID int) (usr types.User, err error) {

	//find user availablity
	//var usr types.User
	tx := dbConn.Begin()
	err = tx.Find(&usr, types.User{UserName: user}).Error
	if err != nil && !strings.EqualFold(err.Error(), "record not found") {
		tx.Rollback()
		Logger.Error(err.Error())
		return
	}
	Logger.Debugf("find user detail %+v", usr)

	//if user not available , create new user
	if usr.UserName == "" {
		err = tx.Save(&types.User{UserName: user, Team: team}).Error
		Logger.Debugf("New user added %s ", user)
		if err != nil {
			tx.Rollback()
			Logger.Error("Saving user failed ", err.Error())
			return
		}
	}

	//TODO : Dirty fix again
	//since it
	r := new(types.UserRole)
	err = tx.Find(&r, types.UserRole{UserID: usr.ID}).Error
	Logger.Debugf("user role details %+v ", r)

	if r.UserRoleID == 0 {
		r.UserID = usr.ID
	}
	//update roles
	r.RoleID = roleID
	tx.Save(&r)
	tx.Commit()
	tx.Close()

	Logger.Debugf("Roles added for the user %s ", user)
	return
}
