package data

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id            uint64         `gorm:"primaryKey, column:id"` // 用户id
	Nickname      string         `gorm:"column:nickname"`       // 用户昵称
	PasswdHash    string         `gorm:"column:passwd_hash"`    // 用户密码hash
	PasswdSalt    string         `gorm:"column:passwd_salt"`    // 用户密码盐
	Fullname      string         `gorm:"column:fullname"`       // 用户全名
	Avatar        string         `gorm:"column:avatar"`         // 用户头像链接
	Email         string         `gorm:"column:email"`          // 用户邮箱
	EmailVerified bool           `gorm:"column:email_verified"` // 用户邮箱验证状态
	Mobile        string         `gorm:"column:mobile"`         // 用户手机号
	Config        string         `gorm:"column:config"`         // 用户配置
	LoginedAt     time.Time      `gorm:"column:logined_at"`     // 用户最后登录时间
	CreatedAt     time.Time      `gorm:"column:created_at"`     // 用户创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at"`     // 用户更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`     // 用户删除时间
}

func (User) TableName() string {
	return "u_user"
}
