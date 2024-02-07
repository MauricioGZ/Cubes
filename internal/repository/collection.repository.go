package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/MauricioGZ/Cubes/internal/entity"
)

const (
	qryInsertCubeToCollection = `insert into OWNED_BY (user_id,cube_id,owned_at,quantity) values(?,?,?,?);`
	qryGetAllCubes            = `select CUBES.*, OWNED_BY.owned_at, OWNED_BY.quantity
									from OWNED_BY
									join CUBES on OWNED_BY.cube_id = CUBES.id
									join USERS on OWNED_BY.user_id = USERS.id
									where USERS.email = ?;`
	qryGetCubeFromCollectionByID = `select CUBES.*, OWNED_BY.owned_at, OWNED_BY.quantity
										from OWNED_BY
										join CUBES on OWNED_BY.cube_id = CUBES.id
										join USERS on OWNED_BY.user_id = USERS.id
										where USERS.email = ? and OWNED_BY.cube_id = ?;`
	qryDeleteCubeFromCollectionByID = `delete from OWNED_BY where user_id = ? and cube_id = ?;`
	qryGetPrimeryKyByUserID         = `select user_id, cube_id from OWNED_BY where user_id = ?;`
)

var (
	ErrItemNotFoundInCollection = errors.New("item not found")
)

func (r *repo) SaveItemToCollection(ctx context.Context, userID, cubeID int64) error {
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

func (r *repo) GetItemsFromCollection(ctx context.Context, email string) ([]entity.Collection, error) {
	var cubes []entity.Collection
	var cube entity.Collection

	rows, err := r.db.QueryContext(
		ctx,
		qryGetAllCubes,
		email,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&cube.ID,
			&cube.Name,
			&cube.Brand,
			&cube.Shape,
			&cube.Owned_at,
			&cube.Quantity,
		)
		if err != nil {
			return nil, err
		}
		cubes = append(cubes, cube)
	}

	return cubes, nil
}

func (r *repo) GetFromCollectionByItemID(ctx context.Context, email string, cubeID int64) (*entity.Collection, error) {
	var cube *entity.Collection

	err := r.db.QueryRowContext(
		ctx,
		qryGetCubeFromCollectionByID,
		email,
		cubeID,
	).Scan(
		cube.ID,
		cube.Name,
		cube.Brand,
		cube.Shape,
		cube.Owned_at,
		cube.Quantity,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrItemNotFoundInCollection
		}
		return nil, err
	}

	return cube, nil
}

func (r *repo) DeleteFromCollectionByItemID(ctx context.Context, userID, cubeID int64) error {
	_, err := r.db.ExecContext(ctx, qryDeleteCubeFromCollectionByID, userID, cubeID)
	return err
}

func (r *repo) GetCollectionByUserID(ctx context.Context, userID int64) ([]entity.ColelctionPrimaryKey, error) {
	var keys []entity.ColelctionPrimaryKey
	var key entity.ColelctionPrimaryKey

	rows, err := r.db.QueryContext(
		ctx,
		qryGetPrimeryKyByUserID,
		userID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&key.UserID,
			&key.CubeID,
		)
		if err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}

	return keys, nil
}
