package repository

import (
	"context"
	"database/sql"

	"github.com/MauricioGZ/Cubes/internal/entity"
)

// brief: this interface wraps the basic CRUD operations.
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	DeleteUserByEmail(ctx context.Context, email string) error
	SaveCube(ctx context.Context, name, brand, shape, image string) error
	GetAllCubes(ctx context.Context, email string) ([]entity.Cube, error)
}

type repo struct {
	db *sql.DB
}

func New(_db *sql.DB) Repository {
	return &repo{
		db: _db,
	}
}
