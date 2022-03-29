package handlers

import (
	"pool-event/dao"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func EventCreate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.EventDB)
	if err = vcago.BindAndValidate(c, body); err != nil {
		return
	}
	if err = body.Create(ctx); err != nil {
		return
	}
	return vcago.NewCreated("event", body)
}

func EventGetByID(c echo.Context) (err error) {
	ctx := c.Request().Context()
	result := new(dao.Event)
	if err = result.Get(ctx, bson.M{"_id": c.Param("id")}); err != nil {
		return
	}
	return vcago.NewSelected("event", result)
}

func EventList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.EventQuery)
	result := new(vcapool.EventList)
	if result, err = body.List(ctx); err != nil {
		return
	}
	return vcago.NewSelected("event_list", result)
}

func EventUpdate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.Event)
	if err = body.Update(ctx); err != nil {
		return
	}
	return vcago.NewUpdated("event", body)
}

func EventDeleteByID(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.Event)
	id := c.Param("id")
	if err = body.Delete(ctx, bson.M{"_id": id}); err != nil {
		return
	}
	return vcago.NewDeleted("event", id)
}
