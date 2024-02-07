package entity

import "time"

type Collection struct {
	Cube
	Owned_at time.Time
	Quantity int64
}

type ColelctionPrimaryKey struct {
	UserID int64
	CubeID int64
}
