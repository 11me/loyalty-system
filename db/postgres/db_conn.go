package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConn struct {
	*sqlx.DB
}

func NewDB(dsn string) (*DBConn, error) {
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}
	return &DBConn{conn}, nil
}
