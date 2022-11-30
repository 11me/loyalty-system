package db

import (
	"context"
	"loyalty-system/model"
)

type User interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUserByID(ctx context.Context, id int) error
}
