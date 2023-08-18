package utils

import (
	"time"

	"github.com/bryanck29/be-test/internal/config"
	"github.com/bryanck29/be-test/internal/schema/model"

	"github.com/dgrijalva/jwt-go"
)

func CreateJwt(user model.User, expiryDuration time.Duration) (string, error) {
	var err error
	timeNow := time.Now().UTC()
	claims := model.JwtClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: timeNow.Add(expiryDuration).Unix(),
			IssuedAt:  timeNow.Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(config.Core.SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token, secret string) (t *jwt.Token, err error) {
	jwtToken, err := jwt.Parse(
		token,
		func(t *jwt.Token) (i interface{}, err error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return
			}

			i = []byte(secret)

			return
		},
	)
	if err != nil {
		return
	}

	t = jwtToken

	return
}
