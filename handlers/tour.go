package handlers

import (
	"pool-event/dao"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type TourHandler struct {
	vcago.Handler
}

func NewTourHandler() *TourHandler {
	handler := vcago.NewHandler("tour")
	return &TourHandler{
		*handler,
	}
}

func (i *TourHandler) Create(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.TourCreate)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	token := new(vcapool.AccessToken)
	if err = c.AccessToken(token); err != nil {
		return
	}
	result := new(vcapool.Tour)
	if result, err = body.Create(c.Ctx(), token); err != nil {
		return
	}
	return c.Created(result)
}

func (i *TourHandler) GetByID(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	result := new(dao.Tour)
	if err = result.Get(c.Ctx(), bson.M{"_id": c.Param("id")}); err != nil {
		return
	}
	return c.Selected(result)
}

func (i *TourHandler) List(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.TourQuery)
	result := new(vcapool.TourList)
	if result, err = body.List(c.Ctx()); err != nil {
		return
	}
	return c.Listed(result)
}

func (i *TourHandler) Update(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.TourUpdate)
	result := new(vcapool.Tour)
	if result, err = body.Update(c.Ctx()); err != nil {
		return
	}
	return c.Updated(result)
}

func (i *TourHandler) DeleteByID(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.Tour)
	id := c.Param("id")
	if err = body.Delete(c.Ctx(), bson.M{"_id": id}); err != nil {
		return
	}
	return c.Deleted(id)
}
