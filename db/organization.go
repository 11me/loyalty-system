package db

import (
	"context"
	"loyalty-system/model"
)

type Organization interface {
	GetOrganizationByID(ctx context.Context, id int) (*model.Organization, error)
	CreateOrganization(ctx context.Context, organization *model.Organization, userID int) error
	UpdateOrganization(ctx context.Context, user *model.Organization, columns ...string) error
	DeleteOrganizationByID(ctx context.Context, id int) error
}
