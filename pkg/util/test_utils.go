package util

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func SetupDb(ctx context.Context) *sqlx.DB {
	db, err := sqlx.Open("postgres", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		fmt.Println(err)
		panic("Cannot open connection to the database")
	}

	return db
}

func EmptyTable(ctx context.Context, db *sqlx.DB, tableName string) {
	_, err := db.ExecContext(ctx, fmt.Sprintf("TRUNCATE TABLE %s", tableName))
	if err != nil {
		log.Fatalf("Cannot clean up table before running tests: %s", err)
	}
}
