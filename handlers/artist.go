package handlers

import (
	"pool-event/dao"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type ArtistHandler struct {
	vcago.Handler
}

func NewArtistHandler() *ArtistHandler {
	handler := vcago.NewHandler("artist")
	return &ArtistHandler{
		*handler,
	}
}

func (i *ArtistHandler) Create(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.Artist)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	if err = body.Create(c.Ctx()); err != nil {
		return
	}
	return c.Created(body)
}

func (i *ArtistHandler) GetByID(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	result := new(dao.Artist)
	if err = result.Get(c.Ctx(), bson.M{"_id": c.Param("id")}); err != nil {
		return
	}
	return c.Selected(result)
}

func (i *ArtistHandler) Update(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.Artist)
	if err = body.Update(c.Ctx()); err != nil {
		return
	}
	return c.Updated(body)
}

func (i *ArtistHandler) DeleteByID(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.Artist)
	id := c.Param("id")
	if err = body.Delete(c.Ctx(), bson.M{"_id": id}); err != nil {
		return
	}
	return c.Deleted(id)
}

func (i *ArtistHandler) List(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.ArtistQuery)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	result := new(vcapool.ArtistList)
	if result, err = body.List(c.Ctx()); err != nil {
		return
	}
	return c.Listed(result)
}
