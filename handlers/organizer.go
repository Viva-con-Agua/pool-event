package handlers

import (
	"pool-event/dao"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type OrganizerHandler struct {
	vcago.Handler
}

func NewOrganizerHandler() *OrganizerHandler {
	handler := vcago.NewHandler("organizer")
	return &OrganizerHandler{
		*handler,
	}
}

func (i *OrganizerHandler) Create(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.Organizer)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	if err = body.Create(c.Ctx()); err != nil {
		return
	}
	return c.Created(body)
}

func (i *OrganizerHandler) GetByID(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	result := new(dao.Organizer)
	if err = result.Get(c.Ctx(), bson.M{"_id": c.Param("id")}); err != nil {
		return
	}
	return c.Selected(result)
}

func (i *OrganizerHandler) Update(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.Organizer)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	if err = body.Update(c.Ctx()); err != nil {
		return
	}
	return c.Updated(body)
}

func (i *OrganizerHandler) DeleteByID(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.Organizer)
	id := c.Param("id")
	if err = body.Delete(c.Ctx(), bson.M{"_id": id}); err != nil {
		return
	}
	return c.Deleted(id)
}

func (i *OrganizerHandler) List(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.OrganizerQuery)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	result := new(vcapool.OrganizerList)
	if result, err = body.List(c.Ctx()); err != nil {
		return
	}
	return c.Listed(result)
}
