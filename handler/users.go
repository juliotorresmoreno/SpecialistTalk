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

func (that *UsersHandler) findUsers(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*model.User)

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

	query := builder.Like{"name", q}.Or(builder.Like{"lastname", q})
	err = conn.
		Where("id <> ?", session.ID).
		And(query).
		Limit(limit, skip).
		Find(&users)
	if err != nil {
		return echo.NewHTTPError(501, helper.ParseError(err).Error())
	}
	return c.JSON(200, users)
}

func (that *UsersHandler) findUser(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*model.User)
	if !session.ACL.IsAdmin() && strconv.Itoa(session.ID) != c.Param("user_id") {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	conf := configs.GetConfig()
	conn, err := db.GetConnectionPoolWithSession(conf.Database, session)
	if err != nil {
		return err
	}

	u := &model.User{}
	_, err = conn.SessionWithACL().Where("id = ?", c.Param("user_id")).Get(u)
	if err != nil {
		return echo.NewHTTPError(501, helper.ParseError(err).Error())
	}
	return c.JSON(200, u)
}

func (that *UsersHandler) addUser(c echo.Context) error {
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(406, helper.ParseError(err).Error())
	}
	conn, err := db.GetConnectionPool()
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}

	if err := u.Check(); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	if _, err := conn.InsertOne(u); err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	return c.JSON(202, u)
}

func (that *UsersHandler) updateUser(c echo.Context) error {
	_session := c.Get("session")
	if _session == nil {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	session := _session.(*model.User)
	if !session.ACL.IsAdmin() && strconv.Itoa(session.ID) != c.Param("user_id") {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	actualUser := &model.User{}
	updateUser := &model.User{}
	if err := c.Bind(updateUser); err != nil {
		return echo.NewHTTPError(406, err.Error())
	}
	conf := configs.GetConfig()
	conn, err := db.GetConnectionPoolWithSession(conf.Database, session)
	if err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}

	_, err = conn.Get(actualUser)
	if actualUser.ID == 0 || err != nil {
		return echo.NewHTTPError(401, "unautorized")
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
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	if _, err := conn.Where("id = ?", actualUser.ID).Update(actualUser); err != nil {
		return echo.NewHTTPError(500, helper.ParseError(err).Error())
	}
	return c.String(http.StatusNoContent, "")
}

// AttachUsers s
func AttachUsers(g *echo.Group) {
	u := &UsersHandler{}
	g.GET("", u.findUsers)
	g.GET("/:user_id", u.findUser)
	g.PUT("", u.addUser)
	g.PATCH("/:user_id", u.updateUser)
}
