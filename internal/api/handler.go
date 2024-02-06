package api

import (
	"log"
	"net/http"

	"github.com/MauricioGZ/Cubes/encryption"
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

func (a *API) GetCubes(c echo.Context) error {
	ctx := c.Request().Context()

	email, err := validateUser(c)
	if err != nil {
		return err
	}

	cubes, err := a.serv.GetOwnedCubes(ctx, email)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			return c.JSON(http.StatusConflict, responseMessage{Message: "Invalid credentials"})
		}
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	return c.JSON(http.StatusOK, cubes)
}

func (a *API) AddCube(c echo.Context) error {
	ctx := c.Request().Context()

	uParams := dtos.AddCube{}
	if err := c.Bind(&uParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err := a.serv.AddCube(
		ctx,
		uParams.Name,
		uParams.Brand,
		uParams.Shape,
		uParams.Image,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	return c.JSON(http.StatusOK, nil)
}

func (a *API) DeleteCube(c echo.Context) error {
	ctx := c.Request().Context()

	email, err := validateUser(c)
	if err != nil {
		return err
	}

	cubes, err := a.serv.GetOwnedCubes(ctx, email)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			return c.JSON(http.StatusConflict, responseMessage{Message: "Invalid credentials"})
		}
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	uParams := dtos.DeleteCube{}
	if err := c.Bind(&uParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	deletedCube := false
	for _, cube := range cubes {
		if uParams.ID == cube.ID {
			err = a.serv.DeleteCube(ctx, uParams.ID)
			deletedCube = true
			break
		}
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	if !deletedCube {
		return c.JSON(http.StatusNotFound, responseMessage{Message: "Cube not found"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "Cube deleted"})
}

func (a *API) AddCubeToCollection(c echo.Context) error {
	ctx := c.Request().Context()

	email, err := validateUser(c)
	if err != nil {
		return err
	}

	cParams := dtos.AddCubeToCollection{}
	if err := c.Bind(&cParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.serv.AddCubeToCollection(ctx, email, cParams.CubeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	return c.JSON(http.StatusOK, nil)
}

func validateUser(c echo.Context) (string, error) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		log.Println(err)
		return "", c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	claims, err := encryption.ParseLoginJWT(cookie.Value)
	if err != nil {
		log.Println(err)
		return "", c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	email := claims["email"].(string)

	return email, nil
}
