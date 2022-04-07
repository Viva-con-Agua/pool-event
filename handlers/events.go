package handlers

import (
	"pool-event/dao"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type EventHandler struct {
	vcago.Handler
}

func NewEventHandler() *EventHandler {
	handler := vcago.NewHandler("event")
	return &EventHandler{
		*handler,
	}
}

func (i *EventHandler) Create(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.EventCreate)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	token := new(vcapool.AccessToken)
	if err = c.AccessToken(token); err != nil {
		return
	}
	result := new(dao.Event)
	if result, err = body.Create(c.Ctx(), token); err != nil {
		return
	}
	return c.Created(result)
}

func (i *EventHandler) GetByID(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	result := new(dao.Event)
	if err = result.Get(c.Ctx(), bson.M{"_id": c.Param("id")}); err != nil {
		return
	}
	return c.Selected(result)
}

func (i *EventHandler) List(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.EventQuery)
	result := new(vcapool.EventList)
	if result, err = body.List(c.Ctx()); err != nil {
		return
	}
	return c.Listed(result)
}

func (i *EventHandler) Update(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.Event)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	if err = body.Update(c.Ctx()); err != nil {
		return
	}
	return c.Updated(body)
}

func (i *EventHandler) DeleteByID(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.Event)
	id := c.Param("id")
	if err = body.Delete(ctx, bson.M{"_id": id}); err != nil {
		return
	}
	return vcago.NewDeleted("event", id)
}
