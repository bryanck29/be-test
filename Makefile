DB_URI ?= $(shell grep -Po '"database_uri": *\K"[^"]*"' config.json | sed 's/"//g')

.PHONY: migrateup migratedown

build:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/echo-swagger
	rm -rf docs
	swag init
	swag fmt
	go build -o backend-test

migrate-up: 
	migrate -path migration -database "$(DB_URI)" up

migrate-down:
	migrate -path migration -database "$(DB_URI)" -verbose down

server:
	rm -rf docs
	swag init
	swag fmt
	go run main.go

test:
	go install github.com/stretchr/gorc && gorc