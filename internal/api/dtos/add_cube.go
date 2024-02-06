package dtos

type AddCube struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Shape string `json:"shape"`
	Image string `json:"image"`
}

type AddCubeToCollection struct {
	CubeID int64 `json:"cube_id"`
}
