package model

import (
	"encoding/json"
	"time"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

// ACL s
type ACL struct {
	Owner string `json:"owner"`
}

type Rol string

const RolUser = "user"
const RolAdmin = "admin"

// User s
type User struct {
	ID            int       `xorm:"id BIGSERIAL not null autoincr pk"         valid:""`
	Username      string    `xorm:"username varchar(20) not null unique"      valid:"username,required"`
	Email         string    `xorm:"email varchar(200) not null unique"        valid:"email,required"`
	Name          string    `xorm:"name varchar(50) not null"                 valid:"name,required"`
	LastName      string    `xorm:"lastname varchar(50) not null"             valid:"name,required"`
	DateBirth     time.Time `xorm:"date_birth DATE"                           valid:""`
	ImgSrc        string    `xorm:"imgSrc text"                               valid:""`
	Country       string    `xorm:"country varchar(2)"                        valid:""`
	Nationality   string    `xorm:"nationality varchar(2)"                    valid:""`
	Facebook      string    `xorm:"facebook varchar(255)"                     valid:""`
	Linkedin      string    `xorm:"linkedin varchar(255)"                     valid:""`
	Password      string    `xorm:"password varchar(100) not null"            valid:""            `
	ValidPassword string    `xorm:"-"                                         valid:"password"    json:"-"`
	RecoveryToken string    `xorm:"recovery_token varchar(100) not null"      valid:""            json:"-"`
	ACL           *ACL      `xorm:"acl json not null"                         valid:"required"    json:"-"`
	CreatedAt     time.Time `xorm:"created_at created"                        valid:""            json:"-"`
	UpdatedAt     time.Time `xorm:"updated_at updated"                        valid:""            json:"-"`
	Version       int       `xorm:"version version"                           valid:""            json:"-"`
}

// TableName s
func (u *User) TableName() string {
	return "users"
}

type user struct {
	ID       int    `json:"id"`
	ACL      ACL    `json:"acl"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`

	DateBirth   time.Time `json:"date_birth"`
	ImgSrc      string    `json:"imgSrc"`
	Country     string    `json:"country"`
	Nationality string    `json:"nationality"`
	Facebook    string    `json:"facebook"`
	Linkedin    string    `json:"linkedin"`

	Password string `json:"password"`
}

type userWithowPassword struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`

	DateBirth   time.Time `json:"date_birth"`
	ImgSrc      string    `json:"imgSrc"`
	Country     string    `json:"country"`
	Nationality string    `json:"nationality"`
	Facebook    string    `json:"facebook"`
	Linkedin    string    `json:"linkedin"`
}

// Check s
func (u *User) Check() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

// SetPassword s
func (u *User) SetPassword(password string) error {
	s, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	u.Password = string(s)
	return nil
}

// UnmarshalJSON s
func (that *User) UnmarshalJSON(b []byte) error {
	u := &user{}
	err := json.Unmarshal(b, u)
	if err != nil {
		return err
	}
	that.ID = 0
	that.Username = u.Username
	that.Email = u.Email
	that.Name = u.Name
	that.LastName = u.LastName
	that.ValidPassword = u.Password

	that.DateBirth = u.DateBirth
	that.ImgSrc = u.ImgSrc
	that.Country = u.Country
	that.Nationality = u.Nationality
	that.Facebook = u.Facebook
	that.Linkedin = u.Linkedin

	that.ACL = &ACL{Owner: u.Username}
	if err != nil {
		return err
	}
	err = that.SetPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}

// MarshalJSON s
func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(userWithowPassword{
		ID:          u.ID,
		Email:       u.Email,
		Username:    u.Username,
		Name:        u.Name,
		LastName:    u.LastName,
		DateBirth:   u.DateBirth,
		ImgSrc:      u.ImgSrc,
		Country:     u.Country,
		Nationality: u.Nationality,
		Facebook:    u.Facebook,
		Linkedin:    u.Linkedin,
	})
}
