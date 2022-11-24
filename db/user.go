package db

import (
	"context"
	"loyalty-system/model"
)

type User interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User, columns ...string) error
	DeleteUserByID(ctx context.Context, id int) error
}
