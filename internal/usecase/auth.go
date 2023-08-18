package usecase

import (
	"encoding/json"

	"github.com/bryanck29/be-test/internal/config"
	"github.com/bryanck29/be-test/internal/constant"
	"github.com/bryanck29/be-test/internal/contract"
	"github.com/bryanck29/be-test/internal/schema/model"
	"github.com/bryanck29/be-test/internal/schema/request"
	"github.com/bryanck29/be-test/internal/schema/response"
	"github.com/bryanck29/be-test/internal/utils"
	extUtils "github.com/bryanck29/be-test/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// authUsecase represents auth usecase object
type authUsecase struct {
	userRepository contract.UserRepository
}

// newAuthUsecase used to intiate auth usecase
func newAuthUsecase(userRepo contract.UserRepository) contract.AuthUsecase {
	return &authUsecase{
		userRepository: userRepo,
	}
}

// Login handles user authentication process
func (u *authUsecase) Login(ctx echo.Context, req request.PostLogin) (result response.PostLogin, err error) {
	user, err := u.userRepository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return
	}

	decryptedPwd, err := extUtils.Decrypt(user.Password, config.Core.SecretKey)
	if err != nil {
		return
	}

	if decryptedPwd != req.Password {
		err = constant.ErrInvalidLoginCredential
		return
	}

	token, err := utils.CreateJwt(user, constant.TokenDefaultDuration)
	if err != nil {
		return
	}

	refreshToken, err := utils.CreateJwt(user, constant.RefreshTokenDefaultDuration)
	if err != nil {
		return
	}

	result = response.PostLogin{
		Token:        token,
		RefreshToken: refreshToken,
	}

	return
}

// RefreshLogin refresh the passed session token
func (u *authUsecase) RefreshLogin(ctx echo.Context, req request.PostRefreshLogin) (result response.PostRefreshLogin, err error) {
	token, err := utils.ParseToken(req.Token, config.Core.SecretKey)
	if err != nil {
		err = constant.ErrInvalidAuth
		return
	}

	if !token.Valid {
		err = constant.ErrInvalidAuthToken
		return
	}

	rawClaims := token.Claims.(jwt.MapClaims)
	jsonBody, err := json.Marshal(rawClaims)
	if err != nil {
		return
	}

	claims := model.JwtClaims{}
	if err = json.Unmarshal(jsonBody, &claims); err != nil {
		return
	}

	newToken, err := utils.CreateJwt(claims.User, constant.TokenDefaultDuration)
	if err != nil {
		return
	}

	newRefreshToken, err := utils.CreateJwt(claims.User, constant.RefreshTokenDefaultDuration)
	if err != nil {
		return
	}

	result = response.PostRefreshLogin{
		PostLogin: response.PostLogin{
			Token:        newToken,
			RefreshToken: newRefreshToken,
		},
	}

	return
}
