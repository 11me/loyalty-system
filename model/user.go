package model

import "time"

type User struct {
	ID           int       `bun:",pk,autoincrement"`
	Name         string    `bun:",notnull"`
	Email        string    `bun:",nullzero,unique"`
	PasswordHash string    `bun:",nullzero,notnull"`
	Admin        *bool     `bun:",nullzero,notnull,default:false"`
	CreatedAt    time.Time `bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
}

type PostUserRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	PasswordPlain string `json:"password"`
}
