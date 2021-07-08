package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// DB represents the database connection.
type DB struct {
	db *sqlx.DB
}

func NewDB(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}