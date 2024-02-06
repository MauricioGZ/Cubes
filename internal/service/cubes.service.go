package service

import (
	"context"

	"github.com/MauricioGZ/Cubes/internal/models"
)

func (s *serv) AddCube(ctx context.Context, name, brand, shape, image string) error {
	return s.repo.SaveCube(ctx, name, brand, shape, image)
}

func (s *serv) DeleteCube(ctx context.Context, id int64) error {
	return s.repo.DeleteCubeByID(ctx, id)
}

func (s *serv) GetOwnedCubes(ctx context.Context, email string) ([]models.Cube, error) {
	cubes, err := s.repo.GetAllCubes(ctx, email)
	var cubesModel []models.Cube

	for _, cube := range cubes {
		cubesModel = append(cubesModel, models.Cube{
			ID:    cube.ID,
			Name:  cube.Name,
			Brand: cube.Brand,
			Shape: cube.Shape,
		})
	}

	return cubesModel, err
}
