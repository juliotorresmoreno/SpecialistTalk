package model

import "time"

type Group struct {
	ID        int       `xorm:"id BIGSERIAL not null autoincr pk"          valid:""`
	Name      string    `xorm:"name varchar(50) not null"                  valid:"required,name"`
	Code      string    `xorm:"code varchar(50) not null"                  valid:"required"`
	Owner     string    `xorm:"owner varchar(100) not null not null index" valid:"required"`
	CreatedAt time.Time `xorm:"created_at created"`
	UpdatedAt time.Time `xorm:"updated_at updated"`
	Version   int       `xorm:"bigint version"`
}

// TableName s
func (u *Group) TableName() string {
	return "groups"
}
