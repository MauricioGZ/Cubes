package api

import (
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterRoutes(e *echo.Echo) {
	user := e.Group("/user")
	cubes := e.Group("/cube")
	collection := e.Group("/collection")

	user.GET("/register", func(c echo.Context) error {
		return c.File("public/signup.html")
	})
	user.POST("/register", a.RegisterUser)
	user.GET("/login", func(c echo.Context) error {
		return c.File("public/login.html")
	})
	user.POST("/login", a.LoginUser)

	cubes.GET("/add", func(c echo.Context) error {
		return c.File("public/addCube.html")
	})
	cubes.POST("/add", a.AddCube)
	//cubes.DELETE("/delete", a.DeleteCube) //TODO: refactor this or implement a soft delete to avoid problems

	collection.GET("", a.GetCubesCollection)
	collection.POST("/add", a.AddCubeToCollection)
	collection.DELETE("/remove", a.RemoveCubeFromCollection)
}
