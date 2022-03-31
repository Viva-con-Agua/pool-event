package handlers

import (
	"pool-event/dao"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func TourCreate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.TourCreate)
	if err = vcago.BindAndValidate(c, body); err != nil {
		return
	}
	result := new(vcapool.Tour)
	if result, err = body.Create(ctx); err != nil {
		return
	}
	return vcago.NewCreated("tour", result)
}

func TourGetByID(c echo.Context) (err error) {
	ctx := c.Request().Context()
	result := new(dao.Tour)
	if err = result.Get(ctx, bson.M{"_id": c.Param("id")}); err != nil {
		return
	}
	return vcago.NewSelected("tour", result)
}

func TourList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.TourQuery)
	result := new(vcapool.TourList)
	if result, err = body.List(ctx); err != nil {
		return
	}
	return vcago.NewSelected("tour_list", result)
}

func TourUpdate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.TourUpdate)
	result := new(vcapool.Tour)
	if result, err = body.Update(ctx); err != nil {
		return
	}
	return vcago.NewUpdated("tour", result)
}

func TourDeleteByID(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.Tour)
	id := c.Param("id")
	if err = body.Delete(ctx, bson.M{"_id": id}); err != nil {
		return
	}
	return vcago.NewDeleted("tour", id)
}
