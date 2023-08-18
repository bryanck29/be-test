package db

import (
	"context"

	"github.com/bryanck29/be-test/internal/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConn(databaseUri string) (*mongo.Client, error) {
	dbClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(databaseUri))
	if err != nil {
		return nil, err
	}

	return dbClient, nil
}

func RunDBMigration(client *mongo.Client) {
	driver, err := mongodb.WithInstance(client, &mongodb.Config{DatabaseName: config.Core.DatabaseName})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", config.Core.DatabaseName, driver)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}
