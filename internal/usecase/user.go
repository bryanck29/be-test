package usecase

import (
	"time"

	"github.com/bryanck29/be-test/internal/config"
	"github.com/bryanck29/be-test/internal/constant"
	"github.com/bryanck29/be-test/internal/contract"
	"github.com/bryanck29/be-test/internal/schema/model"
	"github.com/bryanck29/be-test/internal/schema/request"
	"github.com/bryanck29/be-test/pkg/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// userUsecase represents user usecase object
type userUsecase struct {
	userRepository contract.UserRepository
}

// newUserUsecase used to intiate user usecase
func newUserUsecase(userRepo contract.UserRepository) contract.UserUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

// InsertUser handles user creation
func (u *userUsecase) InsertUser(ctx echo.Context, req request.PostInsertUser) (err error) {
	encryptedPwd, err := utils.Encrypt(req.Password, config.Core.SecretKey)
	if err != nil {
		return
	}

	return u.userRepository.InsertUser(ctx, model.User{
		Id:        uuid.New(),
		Name:      req.Name,
		Username:  req.Username,
		Password:  encryptedPwd,
		Role:      constant.USER_ROLE_USER,
		CreatedAt: time.Now().UTC().Unix(),
		UpdatedAt: time.Now().UTC().Unix(),
	})
}

// GetUsers handles getting users data
func (u *userUsecase) GetUsers(ctx echo.Context) (result []model.User, err error) {
	return u.userRepository.GetUsers(ctx)
}

// GetUser handles getting a user data
func (u *userUsecase) GetUser(ctx echo.Context, req request.GetUser) (result model.User, err error) {
	return u.userRepository.GetUser(ctx, req.UserId)
}

// DeleteUser handles getting a user data
func (u *userUsecase) DeleteUser(ctx echo.Context, req request.DeleteUser) (err error) {
	return u.userRepository.DeleteUser(ctx, req.UserId)
}

// UpdateUser handles updating user data
func (u *userUsecase) UpdateUser(ctx echo.Context, req request.PutUser) (result model.User, err error) {
	user, err := u.userRepository.GetUser(ctx, req.UserId)
	if err != nil {
		return
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Password != "" {
		user.Password = req.Password
	}

	if req.Role != "" {
		user.Role = req.Role
	}

	return u.userRepository.UpdateUser(ctx, user)
}
