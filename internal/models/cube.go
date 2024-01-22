package models

type Cube struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Shape string `json:"shape"`
	Image string `json:"image"`
}
