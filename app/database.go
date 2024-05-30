package app

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mnabil1718/go-restful-api/helper"
)

func NewDB() *sql.DB {
	// connection string format: postgres://username:password@localhost:5432/database_name
	db, err := sql.Open("pgx", "postgres://mnabil:Cucibaju123@localhost:5432/go_restful")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db

}
