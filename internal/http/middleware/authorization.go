package middleware

import (
	"encoding/json"
	"strings"

	"github.com/bryanck29/be-test/internal/config"
	"github.com/bryanck29/be-test/internal/constant"
	"github.com/bryanck29/be-test/internal/schema/model"
	"github.com/bryanck29/be-test/internal/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var AuthUser = Auth(constant.USER_ROLE_USER)
var AuthAdmin = Auth(constant.USER_ROLE_ADMIN)
var AuthUserOrAdmin = Auth("")

// Auth is a middleware that checks bearer token authorization and the role the token owner has
func Auth(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			bearer := c.Request().Header.Get("Authorization")

			token, err := extractBearer(bearer)
			if err != nil {
				err = constant.ErrInvalidAuth
				return utils.ErrorResponse(c, err, "", nil)
			}

			claims, err := verifyToken(c, token)
			if err != nil {
				err = constant.ErrInvalidAuth
				return utils.ErrorResponse(c, err, "", nil)
			}

			buff, err := json.Marshal(claims)
			if err != nil {
				err = constant.ErrInvalidAuth
				return utils.ErrorResponse(c, err, "", nil)
			}

			jwtClaims := model.JwtClaims{}
			json.Unmarshal(buff, &jwtClaims)

			if jwtClaims.User == (model.User{}) {
				err = constant.ErrInvalidSession
				return utils.ErrorResponse(c, err, "", nil)
			}

			if role != "" && jwtClaims.User.Role != role {
				err = constant.ErrInvalidAccess
				return utils.ErrorResponse(c, err, "", nil)
			}
			c.Set(constant.ClaimsUser, jwtClaims.User)

			return next(c)
		}
	}
}

// extractBearer extracts token from bearer
func extractBearer(bearer string) (token string, err error) {
	if bearer == "" {
		err = constant.ErrInvalidAuth
		return
	}

	splittedBearer := strings.Split(bearer, " ")
	if len(splittedBearer) != 2 {
		err = constant.ErrInvalidAuth
		return
	}

	token = splittedBearer[1]

	return
}

// verifyToken used to verify token and returns claims
func verifyToken(c echo.Context, token string) (claims jwt.MapClaims, err error) {
	jwtToken, err := jwt.Parse(
		token,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, constant.ErrInvalidAuth
			}
			return []byte(config.Core.SecretKey), nil
		},
	)
	if err != nil {
		err = constant.ErrInvalidAuth
		return
	}

	claims = jwtToken.Claims.(jwt.MapClaims)
	return
}
