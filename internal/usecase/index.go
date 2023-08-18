package usecase

import (
	"github.com/bryanck29/be-test/internal/contract"
	"github.com/bryanck29/be-test/internal/repository"
)

type AppUsecases struct {
	AuthUsecase contract.AuthUsecase
	UserUsecase contract.UserUsecase
}

// InitUsecases will initialize the app usecases/services
func InitUsecases(repositories repository.AppRepositories) AppUsecases {
	authUsecase := newAuthUsecase(repositories.UserRepository)
	userUsecase := newUserUsecase(repositories.UserRepository)
	return AppUsecases{
		AuthUsecase: authUsecase,
		UserUsecase: userUsecase,
	}
}
