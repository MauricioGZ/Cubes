package api

import (
	"fmt"

	"github.com/MauricioGZ/Cubes/internal/service"
	"github.com/labstack/echo/v4"
)

type API struct {
	serv service.Service
}

func New(_serv service.Service) *API {
	return &API{
		serv: _serv,
	}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)
	return e.Start(fmt.Sprintf(":%s", address))
}
