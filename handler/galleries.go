package handler

import (
	"github.com/juliotorresmoreno/SpecialistTalk/db"
	"github.com/juliotorresmoreno/SpecialistTalk/helper"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/labstack/echo/v4"
)

type GalleryHandler struct {
}

func AttachGallery(g *echo.Group) {
	u := &GalleryHandler{}

	g.GET("/:id", u.get)
	g.GET("", u.find)
	g.POST("", u.add)
	g.PATCH("/:id", u.update)
	g.DELETE("/:id", u.delete)
}

func (u *GalleryHandler) get(c echo.Context) error {
	session, err := helper.ValidateSession(c)
	if err != nil {
		return err
	}
	crud := db.NewCrud(session)
	gallery := new(model.Gallery)

	if err = crud.Get(gallery); err != nil {
		return err
	}

	return c.JSON(200, gallery)
}

func (u *GalleryHandler) find(c echo.Context) error {
	session, err := helper.ValidateSession(c)
	if err != nil {
		return err
	}
	crud := db.NewCrud(session)

	pagination := helper.Paginate(c)
	opts := model.FindOptions{
		Limit: pagination.Limit,
		Skip:  pagination.Skip,
	}
	if pagination.OrderBy != "" {
		opts.OrderBy = pagination.OrderBy
	}

	galleries := make([]model.Gallery, 0)
	if err = crud.Find(&galleries, opts); err != nil {
		return err
	}

	return c.JSON(200, galleries)
}

type GalleryPOSTAdd struct {
	Name string
}

func (u *GalleryHandler) add(c echo.Context) error {
	payload := &GalleryPOSTAdd{}
	session, err := helper.GetPayload(c, payload)
	if err != nil {
		return err
	}
	crud := db.NewCrud(session)

	gallery := &model.Gallery{
		Name: payload.Name,
	}
	if err = crud.Add(gallery); err != nil {
		return err
	}

	return c.JSON(200, gallery)
}

type GalleryPOSTUpdate struct {
	Name string
}

func (u *GalleryHandler) update(c echo.Context) error {
	payload := &GalleryPOSTUpdate{}
	session, err := helper.GetPayload(c, payload)
	if err != nil {
		return err
	}
	crud := db.NewCrud(session)

	id := c.Param("id")
	gallery := &model.Gallery{
		Name: payload.Name,
	}
	if err = crud.Update(id, gallery); err != nil {
		return err
	}

	return c.JSON(200, gallery)
}

func (u *GalleryHandler) delete(c echo.Context) error {
	session, err := helper.ValidateSession(c)
	if err != nil {
		return err
	}
	crud := db.NewCrud(session)

	id := c.Param("id")
	if err = crud.Delete(id); err != nil {
		return err
	}

	return helper.HTTPStatusNotContent
}
