package entity

import "time"

type OwnedBy struct {
	UserID  int64
	CubeID  int64
	OwnedAt time.Time
}
