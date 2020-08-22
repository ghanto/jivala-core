package platform

import (
	"github.com/jmoiron/sqlx"
)

type PostgresDatabase struct {
	*sqlx.DB
}

func NewPostgresDatabase(databaseURL string) (*PostgresDatabase, error) {
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	return &PostgresDatabase{db}, nil
}
