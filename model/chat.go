package model

import "time"

// Chat s
type Chat struct {
	ID        int       `xorm:"id BIGSERIAL not null autoincr pk" valid:""`
	UserId    int       `xorm:"user_id bigint not null"        valid:""`
	ACL       *ACL      `xorm:"acl json not null"               valid:"required"`
	CreatedAt time.Time `xorm:"created_at created"`
	UpdatedAt time.Time `xorm:"updated_at updated"`
	Version   int       `xorm:"version version"`
}

// TableName s
func (u *Chat) TableName() string {
	return "chats"
}
