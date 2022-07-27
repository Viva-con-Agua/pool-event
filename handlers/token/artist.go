package token

import (
	"pool-event/dao"
	"pool-event/models"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcago/vmdb"
	"github.com/labstack/echo/v4"
)

type ArtistHandler struct {
	vcago.Handler
}

var Artist = &ArtistHandler{*vcago.NewHandler("artist")}

func (i *ArtistHandler) Routes(group *echo.Group) {
	group.Use(i.Context)
	group.POST("", i.Create, cookie)
	group.GET("", i.Get, cookie)
	group.GET("/:id", i.GetByID, cookie)
	group.PUT("", i.Update, cookie)
	group.DELETE("/:id", i.Delete, cookie)
}

func (i *ArtistHandler) Create(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(models.ArtistCreate)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	result := body.Artist()
	if err = dao.ArtistCollection.InsertOne(c.Ctx(), result); err != nil {
		return
	}
	return c.Created(result)
}

func (i *ArtistHandler) GetByID(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(models.ArtistParam)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	result := new(models.Artist)
	if err = dao.ArtistCollection.FindOne(c.Ctx(), body.Filter(), result); err != nil {
		return
	}
	return c.Selected(result)
}

func (i *ArtistHandler) Update(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(models.ArtistUpdate)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	result := new(models.Artist)
	if err = dao.ArtistCollection.UpdateOne(c.Ctx(), body.Filter(), vmdb.NewUpdateSet(body), result); err != nil {
		return
	}
	return c.Updated(body)
}

func (i *ArtistHandler) Delete(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(models.ArtistParam)
	if c.BindAndValidate(body); err != nil {
		return
	}
	if err = dao.ArtistCollection.DeleteOne(c.Ctx(), body.Filter()); err != nil {
		return
	}
	return c.Deleted(body.ID)
}

func (i *ArtistHandler) Get(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(models.ArtistQuery)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	result := new([]models.Artist)
	if err = dao.ArtistCollection.Find(c.Ctx(), body.Filter(), result); err != nil {
		return
	}
	return c.Listed(result)
}
