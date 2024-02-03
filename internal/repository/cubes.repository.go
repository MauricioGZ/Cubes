package repository

import (
	"context"

	"github.com/MauricioGZ/Cubes/internal/entity"
)

const (
	qryInsertCube  = `insert into CUBES (name,brand,shape) values(?,?,?);`
	qryGetAllCubes = `select CUBES.id,CUBES.name,CUBES.brand,CUBES.shape,CUBES.owned_by
						from USERS
						inner join CUBES on USERS.id = CUBES.owned_by
						where USERS.email = ?;`
)

func (r *repo) SaveCube(ctx context.Context, name, brand, shape, image string) error {
	_, err := r.db.ExecContext(ctx, qryInsertCube, name, brand, shape)
	return err
}
func (r *repo) GetAllCubes(ctx context.Context, email string) ([]entity.Cube, error) {
	var cubes []entity.Cube
	var cube entity.Cube
	rows, err := r.db.QueryContext(ctx, qryGetAllCubes, email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&cube.ID, &cube.Name, &cube.Brand, &cube.Shape, &cube.OwnedBy)
		if err != nil {
			return nil, err
		}

		cubes = append(cubes, cube)
	}
	return cubes, nil
}
