package token

import (
	"pool-event/dao"
	"pool-event/models"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcago/vmdb"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/labstack/echo/v4"
)

type OrganizerHandler struct {
	vcago.Handler
}

var Organizer = &OrganizerHandler{*vcago.NewHandler("artist")}

func (i *OrganizerHandler) Routes(group *echo.Group) {
	group.Use(i.Context)
	group.POST("", i.Create, vcapool.AccessCookieConfig())
	group.GET("", i.Get, vcapool.AccessCookieConfig())
	group.GET("/:id", i.GetByID, vcapool.AccessCookieConfig())
	group.PUT("", i.Update, vcapool.AccessCookieConfig())
	group.DELETE("/:id", i.Delete, vcapool.AccessCookieConfig())
}

func (i *OrganizerHandler) Create(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(models.OrganizerCreate)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	result := body.Organizer()
	if err = dao.OrganizerCollection.InsertOne(c.Ctx(), result); err != nil {
		return
	}
	return c.Created(result)
}

func (i *OrganizerHandler) GetByID(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(models.OrganizerParam)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	result := new(models.Organizer)
	if err = dao.OrganizerCollection.FindOne(c.Ctx(), body.Filter(), result); err != nil {
		return
	}
	return c.Selected(result)
}

func (i *OrganizerHandler) Update(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(models.OrganizerUpdate)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	result := new(models.Organizer)
	if err = dao.OrganizerCollection.UpdateOne(c.Ctx(), body.Filter(), vmdb.NewUpdateSet(body), result); err != nil {
		return
	}
	return c.Updated(body)
}

func (i *OrganizerHandler) Delete(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(models.OrganizerParam)
	if c.BindAndValidate(body); err != nil {
		return
	}
	if err = dao.OrganizerCollection.DeleteOne(c.Ctx(), body.Filter()); err != nil {
		return
	}
	return c.Deleted(body.ID)
}

func (i *OrganizerHandler) Get(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(models.OrganizerQuery)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	result := new([]models.Organizer)
	if err = dao.OrganizerCollection.Find(c.Ctx(), body.Filter(), result); err != nil {
		return
	}
	return c.Listed(result)
}
