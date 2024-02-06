package service

import (
	"context"
	"errors"
)

var (
	ErrAddingToCollection = errors.New("error adding cube to the collection")
)

func (s *serv) AddCubeToCollection(ctx context.Context, userEmail string, cubeID int64) error {
	user, err := s.repo.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	err = s.repo.SaveCubeToCollection(ctx, user.ID, cubeID)
	if err != nil {
		return ErrAddingToCollection
	}

	return nil
}
