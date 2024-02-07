package repository

import (
	"context"

	"github.com/MauricioGZ/Cubes/internal/entity"
)

const (
	qryInsertUser        = `insert into USERS (email,name,password,role_id) values(?,?,?,?);`
	qryGetUserByEmail    = `select id,email,name,password,role_id from USERS where email=?;`
	qryDeleteUserByEmail = `delete from USERS where email=?;`
)

const (
	AdminRole    int64 = 1
	SellerRole   int64 = 2
	CustomerRole int64 = 3
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password, CustomerRole)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRowContext(
		ctx,
		qryGetUserByEmail,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.RoleID,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repo) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := r.db.ExecContext(ctx, qryDeleteUserByEmail, email)
	return err
}
