package model

import "time"

type Organization struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type PostOrganizationRequest struct {
	Name string `validate:"required,max=100,excludesall=!@#$%^&*()_+.0x2C?" json:"name"`
}
