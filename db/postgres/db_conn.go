package postgres

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DBConn struct {
	*bun.DB
}

func NewDB(dsn string) (*DBConn, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
	))

	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	return &DBConn{
		bun.NewDB(sqldb, pgdialect.New()),
	}, nil
}
