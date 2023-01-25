package handler

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/juliotorresmoreno/freelive/db"
	"github.com/juliotorresmoreno/freelive/helper"
	"github.com/juliotorresmoreno/freelive/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
}

// POSTLoginPayload s
type POSTSingUpPayload struct {
	Email    string `yaml:"email"   `
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	Name     string `yaml:"name"    `
	LastName string `yaml:"lastname"`
}

func (el *AuthHandler) POSTSingUp(c echo.Context) error {
	conn, err := db.NewEngigne()
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	defer conn.Close()

	p := &POSTSingUpPayload{}
	err = c.Bind(p)
	if err != nil {
		return helper.MakeHTTPError(http.StatusBadRequest, errors.New("body has not valid format"))
	}

	u := &model.User{}
	u.Email = p.Email
	exists, err := conn.Get(u)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	if exists {
		return helper.MakeHTTPError(http.StatusUnauthorized, errors.New("user already exists"))
	}

	u.Name = p.Name
	u.LastName = p.LastName
	u.Username = p.Username
	u.ValidPassword = p.Password
	u.ACL = model.NewACL(p.Username, model.RolUser)

	if err = u.Check(); err != nil {
		return helper.MakeHTTPError(http.StatusBadRequest, err)
	}

	err = u.SetPassword(p.Password)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}

	_, err = conn.Table(u.TableName()).InsertOne(u)

	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}

	return helper.MakeSession(c, u)
}

// POSTLoginPayload s
type POSTLoginPayload struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

func (el *AuthHandler) POSTLogin(c echo.Context) error {
	conn, err := db.NewEngigne()
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer conn.Close()
	p := &POSTLoginPayload{}
	err = c.Bind(p)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	u := &model.User{}
	u.Email = p.Email
	_, err = conn.Get(u)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(p.Password),
	)
	if err != nil {
		return helper.MakeHTTPError(http.StatusUnauthorized, errors.New("password: password is not valid"))
	}

	return helper.MakeSession(c, u)
}

// CredentialsRecovery s
type CredentialsRecovery struct {
	Email string
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (el *AuthHandler) POSTRecovery(c echo.Context) error {
	conn, err := db.NewEngigne()
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	defer conn.Close()
	p := &CredentialsRecovery{}
	err = c.Bind(p)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}

	token := StringWithCharset(40, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz")
	if p.Email == "" {
		return helper.MakeHTTPError(http.StatusNotAcceptable, errors.New("email is required"))
	}
	u := &model.User{RecoveryToken: token}
	q := &model.User{Email: p.Email}

	_, err = conn.Omit("acl").Update(u, q)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	return c.String(204, "")
}

// CredentialsReset s
type CredentialsReset struct {
	Password string
	Token    string
}

func (el *AuthHandler) POSTReset(c echo.Context) error {
	conn, err := db.NewEngigne()
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	defer conn.Close()
	p := &CredentialsReset{}
	err = c.Bind(p)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	if p.Token == "-" || p.Token == "" {
		return helper.MakeHTTPError(http.StatusNotAcceptable, errors.New("token is required"))
	}
	q := &model.User{RecoveryToken: p.Token}
	u := &model.User{}
	_, err = conn.Get(u)
	if err != nil {
		return helper.MakeHTTPError(http.StatusNotAcceptable, err)
	}

	err = u.SetPassword(p.Password)
	if err != nil {
		return helper.MakeHTTPError(http.StatusNotAcceptable, err)
	}
	u.RecoveryToken = "-"
	if err := u.Check(); err != nil {
		return helper.MakeHTTPError(http.StatusNotAcceptable, err)
	}
	_, err = conn.Omit("acl").Update(u, q)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	return c.String(204, "")
}

// AuthHandler s
func AttachAuth(g *echo.Group) {
	c := AuthHandler{}
	g.POST("/sing-up", c.POSTSingUp)
	g.POST("/sing-in", c.POSTLogin)
	g.POST("/recovery", c.POSTRecovery)
	g.POST("/reset", c.POSTReset)
}
