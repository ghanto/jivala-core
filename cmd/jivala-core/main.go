package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	r := gin.Default()

	db, err := initDb()
	defer db.Close()

	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	DefineRoutes(r, db)

	err = r.Run()
	if err != nil {
		log.Fatalf("Error while starting server %v", err)
	}
}

func initDb() (*sqlx.DB, error) {
	sqlDB, err := sql.Open("postgres", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		return nil, err
	}
	db := sqlx.NewDb(sqlDB, "postgres")
	return db, nil
}
