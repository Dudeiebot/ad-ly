package services

import (
	"context"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/dudeiebot/ad-ly/config"
	customizedError "github.com/dudeiebot/ad-ly/errors"
	"github.com/dudeiebot/ad-ly/helpers"
	"github.com/dudeiebot/ad-ly/middlewares"
	"github.com/dudeiebot/ad-ly/models"
	"github.com/dudeiebot/ad-ly/request"
	"github.com/dudeiebot/ad-ly/responses"
)

func GenerateCode(
	payload request.GenerateCode,
	r *http.Request,
) (response responses.LinkResponse, err error, status int) {
	userId := middlewares.GetUser(r.Context()).Id
	userKey := "user_key"
	var link models.Link
	ctx := context.Background()
	urlKey := "url_key_" + payload.CustomCode

	val, err := config.Redis.Incr(ctx, userKey).Result()
	if err != nil {
		return response, helpers.ServerError(err), http.StatusInternalServerError
	}

	if payload.CustomCode == "" {
		payload.CustomCode, err = helpers.GenerateShortCode(uint64(val))
		if err != nil {
			return response, helpers.ServerError(err), http.StatusInternalServerError
		}
	}

	_, err = config.Redis.Get(ctx, urlKey).Result()
	if err != nil && err != redis.Nil {
		return response, helpers.ServerError(err), http.StatusInternalServerError
	}
	if err == nil {
		return responses.LinkResponse{
			Code: payload.CustomCode,
		}, nil, http.StatusOK
	}

	var existingLink models.Link
	if err := config.PostDb.Where("code = ?", payload.CustomCode).First(&existingLink).Error; err == nil {
		return response, customizedError.ErrCodeAlreadyExist, http.StatusConflict
	}

	if err := config.PostDb.Where("url = ? AND user_id = ?", payload.Url, userId).First(&existingLink).Error; err == nil {
		return responses.LinkResponse{
			Code:      existingLink.Code,
			ExpiresAt: existingLink.ExpireAt,
		}, nil, http.StatusOK
	}

	expireDays := 1
	if payload.ExpireAt > 0 {
		expireDays = payload.ExpireAt
	}
	expiresTime := time.Now().Add(time.Hour * 24 * time.Duration(expireDays))

	link = models.Link{
		Code:      payload.CustomCode,
		UserId:    userId,
		Url:       payload.Url,
		ExpireAt:  &expiresTime,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = config.Redis.Set(ctx, urlKey, payload.Url, time.Until(expiresTime)).Err(); err != nil {
		return response, helpers.ServerError(err), http.StatusInternalServerError
	}

	if err = config.PostDb.Create(&link).Error; err != nil {
		return response, helpers.ServerError(err), http.StatusInternalServerError
	}

	return responses.LinkResponse{
		Code:      payload.CustomCode,
		ExpiresAt: &expiresTime,
	}, nil, http.StatusOK
}

func GetUrl(code string) (response responses.UrlResponse, err error, status int) {
	var link models.Link
	urlKey := "url_key_" + code
	ctx := context.Background()

	url, err := config.Redis.Get(ctx, urlKey).Result()
	if err != nil && err != redis.Nil {
		return response, helpers.ServerError(err), http.StatusInternalServerError
	}
	if err == nil {
		return responses.UrlResponse{
			Url: url,
		}, nil, http.StatusTemporaryRedirect
	}

	err = config.PostDb.Where("code = ?", code).First(&link).Error
	if err != nil {
		return response, helpers.ServerError(err), http.StatusInternalServerError
	}
	if link.Expired() {
		return response, customizedError.ErrLinkExpired, http.StatusBadRequest
	}

	return responses.UrlResponse{
		Url: link.Url,
	}, nil, http.StatusTemporaryRedirect
}
