package config

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Core config
)

type config struct {
	ServerPort             int    `json:"server_port"`
	DatabaseName           string `json:"database_name"`
	DatabaseUri            string `json:"database_uri"`
	SecretKey              string `json:"secret_key"`
	DefaultTimeoutInSecond int    `json:"default_timeout_in_second"`
	DB                     *mongo.Database
}
