package helpers

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	"github.com/dudeiebot/ad-ly/config"
)

func GenerateAccessToken(ctx context.Context, userId string) (token string, err error) {
	tokenStore := uuid.New().String()
	tokenExpiry := time.Hour * 24

	err = config.Redis.Set(ctx, "user_auth_"+tokenStore, userId, tokenExpiry).Err()
	if err != nil {
		return "", err
	}

	auth := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(tokenExpiry).Unix(),
		"token": tokenStore,
	})

	token, err = auth.SignedString([]byte(config.AppConfig.AppKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
