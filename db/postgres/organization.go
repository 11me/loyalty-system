package postgres

import (
	"context"
	"loyalty-system/model"
)

func (db *DBConn) CreateOrganization(ctx context.Context, organization *model.Organization, userID int) error {
	query := `WITH org_t AS (INSERT INTO organization(name) VALUES ($1) RETURNING id),
				org_u AS (
					INSERT INTO organization_user(user_id, organization_id)
					SELECT $2, org_t.id FROM org_t
				)
				INSERT INTO admin_organization(user_id, organization_id)
				SELECT $2, org_t.id from org_t`

	_, err := db.ExecContext(ctx, query, organization.Name, userID)
	if err != nil {
		return err
	}
	return nil
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
