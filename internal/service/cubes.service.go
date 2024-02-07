package service

import (
	"context"
)

func (s *serv) AddCube(ctx context.Context, userEmail, name, brand, shape, image string) error {
	user, err := s.repo.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return err
	}
	//TODO: thing in a better way to implement this check
	if user.RoleID > 2 {
		return ErrInvalidCredentials
	}

	return s.repo.SaveCube(ctx, name, brand, shape, image)
}

func (s *serv) DeleteCube(ctx context.Context, userEmail string, cubeID int64) error {
	user, err := s.repo.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	if user.RoleID != 1 {
		return ErrInvalidCredentials
	}
	return s.repo.DeleteCubeByID(ctx, cubeID)
}
