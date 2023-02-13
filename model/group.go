package model

import "time"

type Group struct {
	ID        int       `xorm:"id BIGSERIAL not null autoincr pk" valid:""`
	Name      string    `xorm:"name varchar(50) not null"         valid:"required,name"`
	Code      string    `xorm:"code varchar(50) not null"         valid:"required"`
	ACL       *ACL      `xorm:"acl json not null"                 valid:"required"`
	CreatedAt time.Time `xorm:"created_at created"`
	UpdatedAt time.Time `xorm:"updated_at updated"`
	Version   int       `xorm:"version version"`
}

// TableName s
func (u *Group) TableName() string {
	return "groups"
}
