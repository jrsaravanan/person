package types

import (
	"time"
)

type (

	// LoginUser datatype
	LoginUser struct {
		ID              int       `gorm:"column:id" json:"Id,omitempty"`
		UserName        string    `gorm:"column:username" json:"userName,omitempty"`
		Password        string    `gorm:"column:password" json:"password,omitempty"`
		Domain          string    `sql:"-" json:"domain,omitempty"`
		AuthToken       string    `sql:"-" json:"x-auth-token,omitempty"`
		TokenCreateAt   time.Time `sql:"-" json:"tokenCreateAt,omitempty"`
		LastUpdatedTime time.Time `sql:"-" json:"lastUpdatedTime,omitempty"`
	}

	//User datatype
	User struct {
		ID       int    `sql:"AUTO_INCREMENT" gorm:"column:user_id;primary_key" json:"Id,omitempty"`
		UserName string `gorm:"column:username;primary_key" json:"userName,omitempty"`
		Team     string `gorm:"column:team" json:"team,omitempty"`
		Roles    Role   `sql:"-"`
	}

	// Role given roles
	Role struct {
		ID         int      `sql:"AUTO_INCREMENT" gorm:"column:role_id;primary_key" json:"Id,omitempty"`
		RoleName   string   `gorm:"column:role_name" json:"roleName,omitempty"`
		CreateAt   string   `gorm:"column:created_at" json:"createdAt,omitempty"`
		Permission []string `sql:"-" json:"permission,omitempty"`
	}

	// UserRole mapping table
	UserRole struct {
		UserRoleID int    `sql:"AUTO_INCREMENT" gorm:"column:user_role_id;primary_key" json:"userRoleId,omitempty"`
		UserID     int    `gorm:"column:user_id;primary_key" json:"userId,omitempty"`
		RoleID     int    `gorm:"column:role_id;primary_key" json:"roleId,omitempty"`
		CreatedBy  string `gorm:"column:created_by;primary_key" json:"createdBy,omitempty"`
		//UpdateBy
		//UpdateAt
	}
)

// TableName LoingUser table name
func (s *LoginUser) TableName() string {
	return "t_user"
}

// TableName get login_user
func (u *User) TableName() string {
	return "login_user"
}

// TableName roles
func (r *Role) TableName() string {
	return "t_role"
}

// TableName user roles
func (r *UserRole) TableName() string {
	return "t_user_role"
}
