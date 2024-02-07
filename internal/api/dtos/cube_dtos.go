package dtos

type AddCube struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Shape string `json:"shape"`
	Image string `json:"image"`
}

type DeleteCube struct {
	ID int64 `json:"id"`
}
