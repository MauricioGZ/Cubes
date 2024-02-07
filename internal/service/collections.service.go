package service

import (
	"context"
	"errors"

	"github.com/MauricioGZ/Cubes/internal/models"
)

var (
	ErrAddingToCollection      = errors.New("error adding cube to the collection")
	ErrCubeAlreadyInCollection = errors.New("cube already in collection")
)

func (s *serv) AddCubeToCollection(ctx context.Context, userEmail string, cubeID int64) error {
	user, err := s.repo.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	collection, err := s.repo.GetCollectionByUserID(ctx, user.ID)
	if err != nil {
		return ErrAddingToCollection
	}

	for _, item := range collection {
		if item.CubeID == cubeID {
			return ErrCubeAlreadyInCollection
		}
	}

	err = s.repo.SaveItemToCollection(ctx, user.ID, cubeID)
	if err != nil {
		return ErrAddingToCollection
	}

	return nil
}

func (s *serv) GetOwnedCubes(ctx context.Context, userEmail string) ([]models.Cube, error) {
	var cubesModel []models.Cube

	cubes, err := s.repo.GetItemsFromCollection(ctx, userEmail)
	if err != nil {
		return nil, err
	}

	for _, cube := range cubes {
		cubesModel = append(cubesModel, models.Cube{
			ID:       cube.ID,
			Name:     cube.Name,
			Brand:    cube.Brand,
			Shape:    cube.Shape,
			Owned_at: cube.Owned_at,
		})
	}

	return cubesModel, nil
}

func (s *serv) RemoveCubeFromCollection(ctx context.Context, userEmail string, cubeID int64) error {
	user, err := s.repo.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	return s.repo.DeleteFromCollectionByItemID(ctx, user.ID, cubeID)
}
