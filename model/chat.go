package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Chat struct {
	ID            int       `xorm:"id BIGSERIAL not null autoincr pk"    json:"id"            valid:""`
	ToUserName    string    `xorm:"to_user_name varchar(100) not null"   json:"to_user_name"  valid:"required"`
	Name          string    `xorm:"name varchar(50) not null"            json:"name"          valid:"required,name"`
	Code          string    `xorm:"code varchar(50) not null index"      json:"code"          valid:"required"`
	Status        string    `xorm:"status varchar(50) not null"          json:"status"        valid:"required,chat_status"`
	Notifications int       `xorm:"notifications int not null default 0" json:"notifications" valid:""`
	Owner         string    `xorm:"owner varchar(100) not null index"    json:"-"             valid:"required"`
	CreatedAt     time.Time `xorm:"created_at created"                   json:"-"`
	UpdatedAt     time.Time `xorm:"updated_at updated"                   json:"-"`
	DeletedAt     time.Time `xorm:"deleted_at deleted"                   json:"-"`
	Version       int       `xorm:"bigint version"                       json:"-"`
}

const ChatStatusActive = "active"
const ChatStatusInactive = "inactive"
const ChatStatusCreated = "created"

// TableName s
func (u *Chat) TableName() string {
	return "chats"
}

// Check s
func (u *Chat) Check() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}
