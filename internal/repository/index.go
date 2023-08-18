package repository

import (
	"github.com/bryanck29/be-test/internal/contract"

	"go.mongodb.org/mongo-driver/mongo"
)

type AppRepositories struct {
	UserRepository contract.UserRepository
}

// InitRepositories will initialize the app repositories
func InitRepositories(db *mongo.Database) AppRepositories {
	userRepository := newUserRepository(db)
	return AppRepositories{
		UserRepository: userRepository,
	}
}
