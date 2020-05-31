package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/ohunyk/webserver/domain/model"
	"github.com/ohunyk/webserver/usecase/interactor"
	"net/http"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c echo.Context) error
}

func NewUserController(ui interactor.UserInteractor) UserController{
	return &userController{ui}
}

func (uc *userController) GetUsers(c echo.Context) error {
	var u []*model.User
	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}