package service

import (
	"context"
)

func (s *serv) AddCube(ctx context.Context, name, brand, shape, image string) error {
	return s.repo.SaveCube(ctx, name, brand, shape, image)
}

func (s *serv) DeleteCube(ctx context.Context, id int64) error {
	return s.repo.DeleteCubeByID(ctx, id)
}
