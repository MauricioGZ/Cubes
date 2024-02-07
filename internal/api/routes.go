package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {
	user := e.Group("/user")
	cubes := e.Group("/cube")
	collection := e.Group("/collection")

	user.POST("/register", a.RegisterUser)
	user.POST("/login", a.LoginUser)

	cubes.POST("/add", a.AddCube)
	//cubes.DELETE("/delete", a.DeleteCube) //TODO: refactor this or implement a soft delete to avoid problems

	collection.GET("", a.GetCubesCollection) //TODO: refactor this
	collection.POST("/add", a.AddCubeToCollection)
	collection.DELETE("/remove", a.RemoveCubeFromCollection)
}
