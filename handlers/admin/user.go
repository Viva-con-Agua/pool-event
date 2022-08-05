package admin

import (
	"pool-event/dao"
	"pool-event/models"

	"github.com/Viva-con-Agua/vcago"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	vcago.Handler
}

var User = &UserHandler{*vcago.NewHandler("user")}

func (i *UserHandler) Routes(group *echo.Group) {
	group.Use(i.Context)
	group.GET("", i.Create)
}

func (i *UserHandler) Create(cc echo.Context) (err error) {
	c := cc.(vcago.Context)
	users := []models.User{}
	if users, err = dao.UserGetRequest(); err != nil {
		return c.ErrorResponse(err)
	}
	var result = []error{}
	for n := range users {
		if err = dao.UserCollection.InsertOne(c.Ctx(), users[n]); err != nil {
			result = append(result, err)
		}
	}
	return c.Created(result)
}
