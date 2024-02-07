package models

import "time"

type Cube struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Brand    string    `json:"brand"`
	Shape    string    `json:"shape"`
	Owned_at time.Time `json:"owned_at"`
}
