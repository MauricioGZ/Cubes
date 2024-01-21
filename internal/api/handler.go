package api

import (
	"net/http"

	"github.com/MauricioGZ/Cubes/internal/api/dtos"
	"github.com/MauricioGZ/Cubes/internal/service"
	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	uParams := dtos.RegisterUser{}
	if err := c.Bind(&uParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}
	err := a.serv.RegisterUser(ctx, uParams.Email, uParams.Name, uParams.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "User already exists"})
		}
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}
	return c.JSON(http.StatusCreated, nil)
}

func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	uParams := dtos.RegisterUser{}
	if err := c.Bind(&uParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}
	user, err := a.serv.LoginUser(ctx, uParams.Email, uParams.Password)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			return c.JSON(http.StatusConflict, responseMessage{Message: "Invalid credentials"})
		}
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "Welcome " + user.Name})
}
