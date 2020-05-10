package user

import (
	"github.com/labstack/echo/v4"
	"github.com/ohunyk/webserver/domain"
	"net/http"
)

type ResponseError struct{
	Message string `json:"message"`
}

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(e *echo.Echo, us domain.UserService){
	handler := &UserHandler{
		userService: us,
	}
	e.GET("/users", handler.GetByEmail)
}

func (h *UserHandler) GetByEmail(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := h.userService.GetByEmail(ctx, "")
	if err != nil {
		return c.JSON(500, ResponseError{Message: err.Error()} )
	}
	return c.JSON(http.StatusOK, user)
}