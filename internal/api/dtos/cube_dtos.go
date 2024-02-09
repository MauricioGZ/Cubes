package dtos

type AddCube struct {
	Name  string `form:"name"`
	Brand string `form:"brand"`
	Shape string `form:"shape"`
}

type DeleteCube struct {
	ID int64 `json:"cube_id"`
}
