package repository

import (
	"context"
	"time"
)

const (
	qryInsertCubeToCollection = `insert into OWNED_BY (user_id,cube_id,owned_at,quantity) values(?,?,?,?);`
)

func (r *repo) SaveCubeToCollection(ctx context.Context, userID, cubeID int64) error {
	_, err := r.db.ExecContext(
		ctx,
		qryInsertCubeToCollection,
		userID,
		cubeID,
		time.Now().UTC(),
		1,
	)
	return err
}
