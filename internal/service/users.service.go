package service

import (
	"context"
	"errors"

	"github.com/MauricioGZ/Cubes/encryption"
	"github.com/MauricioGZ/Cubes/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {

	user, _ := s.repo.GetUserByEmail(ctx, email)
	if user != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))

	if err != nil {
		return err
	}

	encryptedPass := encryption.ToBase64(bb)

	s.repo.SaveUser(ctx, email, name, encryptedPass)

	return nil
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(user.Password)
	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
