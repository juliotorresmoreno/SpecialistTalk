package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Chat s
type Chat struct {
	ID        int       `xorm:"id BIGSERIAL not null autoincr pk" valid:""                     json:"id"`
	UserID    int       `xorm:"user_id BIGINT not null"           valid:"required"             json:"user_id"`
	Name      string    `xorm:"name varchar(50) not null"         valid:"required,name"        json:"name"`
	Code      string    `xorm:"code varchar(50) not null index"   valid:"required"             json:"code"`
	Status    string    `xorm:"status varchar(50) not null"       valid:"required,chat_status" json:"status"`
	Owner     string    `xorm:"owner varchar(100) not null"       valid:"required"             json:"-"`
	CreatedAt time.Time `xorm:"created_at created"                                             json:"-"`
	UpdatedAt time.Time `xorm:"updated_at updated"                                             json:"-"`
	DeletedAt time.Time `xorm:"deleted_at deleted"                                             json:"-"`
	Version   int       `xorm:"version version"                                                json:"-"`
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
