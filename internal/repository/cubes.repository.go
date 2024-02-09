package repository

import (
	"context"
)

const (
	qryInsertCube     = `insert into CUBES (name,brand,shape,image) values(?,?,?,?);`
	qryDeleteCubeByID = `delete from CUBES where id = ?;`
)

func (r *repo) SaveCube(ctx context.Context, name, brand, shape, image string) error {
	//TODO: implement the image (AWS for example)
	_, err := r.db.ExecContext(ctx, qryInsertCube, name, brand, shape, image)
	return err
}

func (r *repo) DeleteCubeByID(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, qryDeleteCubeByID, id)
	return err
}
