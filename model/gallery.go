package model

import (
	"time"
)

type Gallery struct {
	ID        int       `xorm:"id BIGSERIAL not null autoincr pk"          valid:""`
	Name      string    `xorm:"name varchar(50) not null"                  valid:"name,maxstringlength(50),required"`
	Owner     string    `xorm:"owner varchar(100) not null not null index" valid:"required"`
	CreatedAt time.Time `xorm:"created_at created"                         json:"-"`
	UpdatedAt time.Time `xorm:"updated_at updated"                         json:"-"`
	DeletedAt time.Time `xorm:"deleted_at deleted default null"            json:"-"`
	Version   int       `xorm:"bigint version"                             json:"-"`
}

// TableName s
func (u *Gallery) TableName() string {
	return "galleries"
}

// TableName s
func (u *Gallery) Check() error {
	return nil
}
