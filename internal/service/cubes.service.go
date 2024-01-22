package service

import (
	"context"

	"github.com/MauricioGZ/Cubes/internal/models"
)

func (s *serv) GetOwnedCubes(ctx context.Context, user *models.User) ([]models.Cube, error) {
	if user != nil {
		return nil, ErrInvalidCredentials
	}

	cubes, err := s.repo.GetAllCubes(ctx, user.ID)
	var cubesModel []models.Cube

	for _, cube := range cubes {
		cubesModel = append(cubesModel, models.Cube{
			ID:    cube.ID,
			Name:  cube.Name,
			Brand: cube.Brand,
			Shape: cube.Shape,
			Image: cube.Image,
		})
	}

	return cubesModel, err
}
