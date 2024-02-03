package service

import (
	"context"

	"github.com/MauricioGZ/Cubes/internal/models"
	"github.com/MauricioGZ/Cubes/internal/repository"
)

type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	GetOwnedCubes(ctx context.Context, email string) ([]models.Cube, error)
}

type serv struct {
	repo repository.Repository
}

func New(_repo repository.Repository) Service {
	return &serv{
		repo: _repo,
	}
}
