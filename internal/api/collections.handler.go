package api

import (
	"net/http"

	"github.com/MauricioGZ/Cubes/internal/api/dtos"
	"github.com/MauricioGZ/Cubes/internal/service"
	"github.com/labstack/echo/v4"
)

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

func (a *API) GetCubesCollection(c echo.Context) error {
	ctx := c.Request().Context()

	email, err := validateUser(c)
	if err != nil {
		return err
	}

	cubes, err := a.serv.GetOwnedCubes(ctx, email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	return c.JSON(http.StatusOK, cubes)
}
