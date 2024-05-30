package app

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mnabil1718/go-restful-api/helper"
)

type DBEnv int

// enums passed in as DB env
const (
	Dev DBEnv = iota
	Test
)

func NewDB(env DBEnv) *sql.DB {
	// connection string format: postgres://username:password@localhost:5432/database_name
	var connString string = "postgres://mnabil:Cucibaju123@localhost:5432/go_restful" // DEV

	if env == Test {
		connString = "postgres://mnabil:Cucibaju123@localhost:5432/go_restful_test" // TEST
	}

	db, err := sql.Open("pgx", connString)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
