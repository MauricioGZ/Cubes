package repository

import (
	"context"

	"github.com/MauricioGZ/Cubes/internal/entity"
)

const (
	qryInsertCube  = `insert into CUBES (name,brand,shape) values(?,?,?);`
	qryGetAllCubes = `select id,name,brand,shape from CUBES where owned_by=?;`
)

func (r *repo) SaveCube(ctx context.Context, name, brand, shape, image string) error {
	_, err := r.db.ExecContext(ctx, qryInsertCube, name, brand, shape)
	return err
}
func (r *repo) GetAllCubes(ctx context.Context, owned_by int64) ([]entity.Cube, error) {
	var cubes []entity.Cube
	var cube entity.Cube
	rows, err := r.db.QueryContext(ctx, qryGetAllCubes, owned_by)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&cube.ID, &cube.Name, &cube.Brand, &cube.Shape, &cube.Image)
		if err != nil {
			return nil, err
		}
		cubes = append(cubes, cube)
	}
	return cubes, nil
}
