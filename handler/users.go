package handler

import (
	"net/http"
	"strconv"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/db"
	"github.com/juliotorresmoreno/SpecialistTalk/helper"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/labstack/echo/v4"
	"xorm.io/builder"
)

type UsersHandler struct {
}

func (that *UsersHandler) find(c echo.Context) error {
	session, err := helper.ValidateSession(c)
	if session == nil {
		return err
	}

	conn, err := db.GetConnectionPool()
	if err != nil {
		return err
	}
	users := make([]model.User, 0)

	limit, skip := helper.Paginate(c)
	q := c.QueryParam("q")
	if q == "" {
		return c.JSON(200, users)
	}

	query := builder.Like{"lower(name)", q}.Or(builder.Like{"lower(lastname)", q})
	err = conn.
		Where("id <> ?", session.ID).
		And(query).
		Limit(limit, skip).
		Find(&users)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(200, users)
}

func (that *UsersHandler) get(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return helper.HTTPErrorUnauthorized
	}

	conn, err := db.GetConnectionPool()
	if err != nil {
		return err
	}

	u := &model.User{}
	_, err = conn.SessionWithACL().Where("id = ?", c.Param("user_id")).Get(u)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(200, u)
}

func (that *UsersHandler) add(c echo.Context) error {
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		return helper.MakeHTTPError(http.StatusBadRequest, err)
	}
	conn, err := db.GetConnectionPool()
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}

	if err := u.Check(); err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	if _, err := conn.InsertOne(u); err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(202, u)
}

func (that *UsersHandler) update(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return helper.HTTPErrorUnauthorized
	}
	session := _session.(*model.User)
	if strconv.Itoa(session.ID) != c.Param("user_id") {
		return helper.HTTPErrorUnauthorized
	}
	actualUser := &model.User{}
	updateUser := &model.User{}
	if err := c.Bind(updateUser); err != nil {
		return helper.MakeHTTPError(http.StatusBadRequest, err)
	}
	conf := configs.GetConfig()
	conn, err := db.GetConnectionPoolWithSession(conf.Database, session)
	if err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}

	_, err = conn.Get(actualUser)
	if actualUser.ID == 0 || err != nil {
		return helper.HTTPErrorUnauthorized
	}
	actualUser.Password = ""
	actualUser.ValidPassword = ""
	actualUser.Name = newValueString(updateUser.Name, actualUser.Name)
	actualUser.LastName = newValueString(updateUser.LastName, actualUser.LastName)

	actualUser.DateBirth = newValueTime(updateUser.DateBirth, actualUser.DateBirth)
	actualUser.ImgSrc = newValueString(updateUser.ImgSrc, actualUser.ImgSrc)
	actualUser.Country = newValueString(updateUser.Country, actualUser.Country)
	actualUser.Nationality = newValueString(updateUser.Nationality, actualUser.Nationality)
	actualUser.Facebook = newValueString(updateUser.Facebook, actualUser.Facebook)
	actualUser.Linkedin = newValueString(updateUser.Linkedin, actualUser.Linkedin)

	if err := actualUser.Check(); err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	if _, err := conn.Where("id = ?", actualUser.ID).Update(actualUser); err != nil {
		return helper.MakeHTTPError(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusNoContent, "")
}

// AttachUsers s
func AttachUsers(g *echo.Group) {
	u := &UsersHandler{}
	g.GET("", u.find)
	g.GET("/:user_id", u.get)
	g.PUT("", u.add)
	g.PATCH("/:user_id", u.update)
}
