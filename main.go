package main

import (
	"log"

	"github.com/MauricioGZ/Cubes/internal/api"
	"github.com/MauricioGZ/Cubes/internal/db"
	"github.com/MauricioGZ/Cubes/internal/repository"
	"github.com/MauricioGZ/Cubes/internal/service"
	"github.com/MauricioGZ/Cubes/settings"
	"github.com/labstack/echo/v4"
)

func main() {

	settings, err := settings.New()
	if err != nil {
		log.Fatal(err)
	}
	db, err := db.New(*settings)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.New(db)
	serv := service.New(repo)
	a := api.New(serv)
	e := echo.New()
	a.Start(e, settings.Port)

	//err = repo.SaveUser(ctx, "correo@corre.com", "nombre", "clave")
	//user, err := repo.GetUserByEmail(ctx, "correo@corre.com")
	//err = repo.DeleteUserByEmail(ctx, "correo@corre.com")
	//err = serv.RegisterUser(ctx, "correo2@correo.com", "Mao", "123")
	//u, err := serv.LoginUser(ctx, "correo@correo.com", "123")
}
