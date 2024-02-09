package api

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/MauricioGZ/Cubes/internal/api/dtos"
	"github.com/labstack/echo/v4"
)

const localStoragePath string = "public/uploads/"

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

	//TODO: add file validation
	//get the file using the method FormFile
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusUnsupportedMediaType, responseMessage{Message: err.Error()})
	}

	//open the received file using the method open and defer the close method
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error while processing media file"})
	}
	defer src.Close()

	//create a new file with the received filename and add the path of the local storage
	dst, err := os.Create(localStoragePath + file.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error while processing media file"})
	}
	defer dst.Close()

	//copy the received file to the local file
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error while processing media file"})
	}

	err = a.serv.AddCube(
		ctx,
		email,
		cParams.Name,
		cParams.Brand,
		cParams.Shape,
		file.Filename, //TODO: thing if this is okay or if it is better to allow to the user to modify the file name
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
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "Cube deleted"})
}

func (a *API) AddCubeImage(c echo.Context) error {
	//ctx := c.Request().Context()
	fmt.Println(("ping"))
	//TODO: add file validation
	//get the file using the method FormFile
	file, err := c.FormFile("image")
	fmt.Println(file.Filename)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnsupportedMediaType, responseMessage{Message: err.Error()})
	}

	//open the received file using the method open and defer the close method
	src, err := file.Open()
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error while processing media file"})
	}
	defer src.Close()

	//create a new file with the received filename and add the path of the local storage
	dst, err := os.Create(localStoragePath + file.Filename)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error while processing media file"})
	}
	defer dst.Close()

	//copy the received file to the local file
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error while processing media file"})
	}
	return c.JSON(http.StatusOK, "ok")
}
