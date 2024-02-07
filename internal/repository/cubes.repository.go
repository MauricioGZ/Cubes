package repository

import (
	"context"
)

const (
	qryInsertCube     = `insert into CUBES (name,brand,shape) values(?,?,?);`
	qryDeleteCubeByID = `delete from CUBES where id = ?;`
)

func (r *repo) SaveCube(ctx context.Context, name, brand, shape, image string) error {
	_, err := r.db.ExecContext(ctx, qryInsertCube, name, brand, shape)
	return err
}

func (r *repo) DeleteCubeByID(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, qryDeleteCubeByID, id)
	return err
}
