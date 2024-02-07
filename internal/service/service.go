package service

import (
	"context"

	"github.com/MauricioGZ/Cubes/internal/models"
	"github.com/MauricioGZ/Cubes/internal/repository"
)

type Service interface {
	RegisterUser(ctx context.Context, userEmail, name, password string) error
	LoginUser(ctx context.Context, userEmail, password string) (*models.User, error)

	AddCube(ctx context.Context, userEmail, name, brand, shape, image string) error
	DeleteCube(ctx context.Context, userEmail string, cubeID int64) error

	AddCubeToCollection(ctx context.Context, userEmail string, cubeID int64) error
	GetOwnedCubes(ctx context.Context, useEmail string) ([]models.Cube, error)
	RemoveCubeFromCollection(ctx context.Context, userEmail string, cubeID int64) error
}

type serv struct {
	repo repository.Repository
}

func New(_repo repository.Repository) Service {
	return &serv{
		repo: _repo,
	}
}
