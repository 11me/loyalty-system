package postgres

import (
	"context"
	"loyalty-system/model"
)

func (db *DBConn) CreateOrganization(ctx context.Context, organization *model.Organization) error {
	query := `INSERT INTO loyalty_system.organization (name) values ($1)`
	_, err := db.Exec(query, organization.Name)
	return err
}

func (db *DBConn) GetOrganizationByID(ctx context.Context, id int) (*model.Organization, error) {
	return nil, nil
}

func (db *DBConn) UpdateOrganization(ctx context.Context, user *model.Organization, columns ...string) error {
	return nil
}

func (db *DBConn) DeleteOrganizationByID(ctx context.Context, id int) error {
	return nil
}
