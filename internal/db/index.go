package db

import (
	"log"

	"github.com/bryanck29/be-test/internal/config"
)

func InitDB() {
	db, err := NewMongoDBConn(config.Core.DatabaseUri)
	if db == nil {
		log.Fatal(err)
	}

	config.Core.DB = db.Database(config.Core.DatabaseName)
	RunDBMigration((db))
}
