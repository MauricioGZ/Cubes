package api

import (
	"fmt"

	"github.com/MauricioGZ/Cubes/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type API struct {
	serv          service.Service
	dataValidator *validator.Validate
}

func New(_serv service.Service) *API {
	return &API{
		serv:          _serv,
		dataValidator: validator.New(),
	}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)
	return e.Start(fmt.Sprintf(":%s", address))
}
