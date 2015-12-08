package types

type (

	// LoginUser datatype
	LoginUser struct {
		ID       int    `gorm:"column:id" json:"Id,omitempty"`
		UserName string `gorm:"column:username" json:"userName,omitempty"`
		Password string `gorm:"column:password" json:"password,omitempty"`
	}

	//User datatype
	User struct {
		ID       int     `sql:"AUTO_INCREMENT" gorm:"column:user_id;primary_key" json:"Id,omitempty"`
		UserName string  `gorm:"column:username;primary_key" json:"userName,omitempty"`
		Team     string  `gorm:"column:team" json:"team,omitempty"`
		Roles    [1]Role `sql:"-"`
	}

	// Role given roles
	Role struct {
		ID       int    `sql:"AUTO_INCREMENT" gorm:"column:role_id;primary_key" json:"Id,omitempty"`
		RoleName string `gorm:"column:role_name" json:"role_Name,omitempty"`
		CreateAt string `gorm:"column:created_at" json:"createdAt,omitempty"`
	}
)

//TableName LoingUser table name
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
