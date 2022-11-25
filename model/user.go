package model

import "time"

type User struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Admin        bool      `db:"admin"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type PostUserRequest struct {
	Name          string `validate:"required" json:"name"`
	Email         string `validate:"required,email" json:"email"`
	PasswordPlain string `validate:"required,gte:8,lte:30" json:"password"`
}
