package model

import "time"

type Group struct {
	ID        int       `xorm:"id BIGSERIAL not null autoincr pk" valid:""`
	Name      string    `xorm:"name varchar(50) not null"         valid:"required,name"`
	Code      string    `xorm:"code varchar(50) not null"         valid:"required"`
	Owner     string    `xorm:"owner varchar(100) not null index" valid:"required"      json:"-" `
	CreatedAt time.Time `xorm:"created_at created"                                      json:"-"`
	UpdatedAt time.Time `xorm:"updated_at updated"                                      json:"-"`
	DeletedAt time.Time `xorm:"deleted_at deleted default null"                         json:"-"`
	Version   int       `xorm:"bigint version"                                          json:"-"`
}

// TableName s
func (u *Group) TableName() string {
	return "groups"
}
