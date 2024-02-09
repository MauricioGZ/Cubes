package api

import (
	"net/http"

	"github.com/MauricioGZ/Cubes/encryption"
	"github.com/MauricioGZ/Cubes/internal/api/dtos"
	"github.com/MauricioGZ/Cubes/internal/service"
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	uParams := dtos.RegisterUser{}

	if err := c.Bind(&uParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err := a.dataValidator.Struct(uParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, uParams.Email, uParams.Name, uParams.Password)
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
	uParams := dtos.LoginUser{}

	if err := c.Bind(&uParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err := a.dataValidator.Struct(uParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	user, err := a.serv.LoginUser(ctx, uParams.Email, uParams.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	token, err := encryption.SignedLoginToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, responseMessage{Message: "Welcome " + user.Name})
}

func validateUser(c echo.Context) (string, error) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		return "", c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	claims, err := encryption.ParseLoginJWT(cookie.Value)
	if err != nil {
		return "", c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	email := claims["email"].(string)

	return email, nil
}
