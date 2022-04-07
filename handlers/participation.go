package handlers

import (
	"pool-event/dao"

	"github.com/Viva-con-Agua/vcago"
	"github.com/Viva-con-Agua/vcapool"
	"github.com/labstack/echo/v4"
)

type ParticipationHandler struct {
	vcago.Handler
}

func NewParticipationHandler() *ParticipationHandler {
	handler := vcago.NewHandler("participation")
	return &ParticipationHandler{
		*handler,
	}
}

func (i *ParticipationHandler) Create(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.ParticipationCreate)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	token := new(vcapool.AccessToken)
	if err = c.AccessToken(token); err != nil {
		return
	}
	result := new(vcapool.Participation)
	if result, err = body.Create(c.Ctx(), token); err != nil {
		return
	}
	return c.Created(result)
}

func (i *ParticipationHandler) Get(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.ParticipationParam)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	token := new(vcapool.AccessToken)
	if err = c.AccessToken(token); err != nil {
		return
	}
	result := new(vcapool.Participation)
	if result, err = body.Get(c.Ctx(), token); err != nil {
		return
	}
	return c.Selected(result)
}

func (i *ParticipationHandler) Update(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.ParticipationUpdate)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	token := new(vcapool.AccessToken)
	if err = c.AccessToken(token); err != nil {
		return
	}
	result := new(vcapool.Participation)
	if result, err = body.Update(c.Ctx(), token); err != nil {
		return
	}
	return c.Updated(result)
}

func (i *ParticipationHandler) List(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.ParticipationQuery)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	token := new(vcapool.AccessToken)
	if err = c.AccessToken(token); err != nil {
		return
	}
	result := new(vcapool.ParticipationList)
	if result, err = body.List(c.Ctx(), token); err != nil {
		return
	}
	return c.Listed(result)
}

func (i *ParticipationHandler) Delete(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	body := new(dao.ParticipationParam)
	if err = c.BindAndValidate(body); err != nil {
		return
	}
	token := new(vcapool.AccessToken)
	if err = c.AccessToken(token); err != nil {
		return
	}
	if err = body.Delete(c.Ctx(), token); err != nil {
		return
	}
	return c.Deleted(body.ID)
}
