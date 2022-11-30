package postgres

import (
	"context"
	"database/sql"
	"loyalty-system/model"
)

func (db *DBConn) CreateUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO "user"
              (first_name, last_name, email, password_hash)
              VALUES ($1, $2, $3, $4)`
	_, err := db.ExecContext(ctx, query, user.FirstName, user.LastName, user.Email, user.PasswordHash)
	return err
}

func (db *DBConn) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}

func (db *DBConn) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM "user" WHERE email = $1`
	err := db.GetContext(ctx, &user, query, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DBConn) UpdateUser(ctx context.Context, user *model.User) error {
	query := `UPDATE user 
              SET first_name      = $1,
                  last_name       = $2,
				  email           = $3,
				  password_hash   = $4
                `
	_, err := db.ExecContext(ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash)
	return err
}

func (db *DBConn) DeleteUserByID(ctx context.Context, id int) error {
	return nil
}
