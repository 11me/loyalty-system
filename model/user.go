package model

import "time"

type User struct {
	ID             int       `db:"id"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	Email          string    `db:"email"`
	PasswordHash   []byte    `db:"password_hash"`
	IsAdmin        bool      `db:"is_admin"`
	OrganizationID *int      `db:"organization_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type PostUserRequest struct {
	FirstName     string `validate:"required,alpha"        json:"first_name"`
	LastName      string `validate:"required,alpha"        json:"last_name"`
	Email         string `validate:"required,email"        json:"email"`
	PasswordPlain string `validate:"required,gte=8,lte=30" json:"password"`
}

type AuthUserRequest struct {
	Email         string `validate:"required,email" json:"email"`
	PasswordPlain string `validate:"required"       json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
