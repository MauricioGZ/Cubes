package api

import (
	"fmt"
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

func (a *API) GetCubesCollection(c echo.Context) error {
	ctx := c.Request().Context()

	email, err := validateUser(c)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cubes, err := a.serv.GetOwnedCubes(ctx, email)
	if err != nil {
		fmt.Println(err)
		if err == service.ErrInvalidCredentials {
			return c.JSON(http.StatusConflict, responseMessage{Message: "Invalid credentials"})
		}
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	return c.JSON(http.StatusOK, cubes)
}

func (a *API) AddCube(c echo.Context) error {
	ctx := c.Request().Context()

	email, err := validateUser(c)
	if err != nil {
		return err
	}

	cParams := dtos.AddCube{}
	if err := c.Bind(&cParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.serv.AddCube(
		ctx,
		email,
		cParams.Name,
		cParams.Brand,
		cParams.Shape,
		cParams.Image,
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

	cParams := dtos.DeleteCube{}
	if err := c.Bind(&cParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.serv.DeleteCube(ctx, email, cParams.ID)
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
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
		if err == service.ErrCubeAlreadyInCollection {
			return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	return c.JSON(http.StatusOK, nil)
}

func (a *API) RemoveCubeFromCollection(c echo.Context) error {
	ctx := c.Request().Context()

	email, err := validateUser(c)
	if err != nil {
		return err
	}

	cParams := dtos.RemoveCubeFromCollection{}
	if err := c.Bind(&cParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.serv.RemoveCubeFromCollection(ctx, email, cParams.CubeID)
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
