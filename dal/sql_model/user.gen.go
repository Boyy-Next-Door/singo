// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sql_model

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID            int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Username      string    `gorm:"column:username;type:varchar(64);not null;uniqueIndex:idx_username,priority:1" json:"username"`
	Password      string    `gorm:"column:password;type:varchar(128);not null" json:"password"`
	Salt          string    `gorm:"column:salt;type:varchar(32);not null" json:"salt"`
	RealName      string    `gorm:"column:real_name;type:varchar(64)" json:"real_name"`
	Mobile        string    `gorm:"column:mobile;type:varchar(20);index:idx_mobile,priority:1" json:"mobile"`
	Email         string    `gorm:"column:email;type:varchar(128)" json:"email"`
	UserType      int32     `gorm:"column:user_type;type:tinyint;not null;comment:1:C 2: 3: 4:" json:"user_type"` // 1:C 2: 3: 4:
	Status        int32     `gorm:"column:status;type:tinyint;not null;default:1;comment:0: 1:" json:"status"`    // 0: 1:
	Avatar        string    `gorm:"column:avatar;type:varchar(255);comment:URL" json:"avatar"`                    // URL
	LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime" json:"last_login_time"`
	CreateTime    time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
