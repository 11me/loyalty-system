package postgres

import (
	"context"
	"loyalty-system/model"
)

func (db *DBConn) CreateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (db *DBConn) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}

func (db *DBConn) UpdateUser(ctx context.Context, user *model.User, columns ...string) error {
	return nil
}

func (db *DBConn) DeleteUserByID(ctx context.Context, id int) error {
	return nil
}
