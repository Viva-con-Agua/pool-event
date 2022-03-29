package handlers

import (
	"pool-event/dao"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func ArtistCreate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.Artist)
	if err = vcago.BindAndValidate(c, body); err != nil {
		return
	}
	if err = body.Create(ctx); err != nil {
		return
	}
	return vcago.NewCreated("artist", body)
}

func ArtistGetByID(c echo.Context) (err error) {
	ctx := c.Request().Context()
	result := new(dao.Artist)
	if err = result.Get(ctx, bson.M{"_id": c.Param("id")}); err != nil {
		return
	}
	return vcago.NewSelected("artist", result)
}

func ArtistUpdate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.Artist)
	if err = body.Update(ctx); err != nil {
		return
	}
	return vcago.NewUpdated("artist", body)
}

func ArtistDeleteByID(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.Artist)
	id := c.Param("id")
	if err = body.Delete(ctx, bson.M{"_id": id}); err != nil {
		return
	}
	return vcago.NewDeleted("artist", id)
}

func ArtistList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	body := new(dao.ArtistQuery)
	if err = vcago.BindAndValidate(c, body); err != nil {
		return
	}
	result := new(vcapool.ArtistList)
	if result, err = body.List(ctx); err != nil {
		return
	}
	return vcago.NewSelected("artist", result)
}
