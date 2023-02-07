package model

import "time"

// Chat s
type Chat struct {
	ID        int       `xorm:"id BIGSERIAL not null autoincr pk" valid:""              json:"id"`
	UserID    int       `xorm:"user_id BIGINT not null"           valid:"required"      json:"user_id"`
	Name      string    `xorm:"name varchar(50) not null"         valid:"required,name" json:"name"`
	Code      string    `xorm:"code varchar(50) not null index"   valid:"required"      json:"code"`
	ACL       *ACL      `xorm:"acl json not null"                 valid:"required"      json:"acl"`
	CreatedAt time.Time `xorm:"created_at created"                                      json:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at updated"                                      json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted_at deleted"                                      json:"deleted_at"`
	Version   int       `xorm:"version version"                                         json:"version"`
}

// TableName s
func (u *Chat) TableName() string {
	return "chats"
}
