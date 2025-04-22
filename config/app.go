package config

import (
	"errors"
	"os"
)

var AppConfig App

type App struct {
	AppName         string
	AppKey          string
	Env             string
	AsynqmonService string
	AppHost         string
}

func loadAppEnv() error {
	appKey, exists := os.LookupEnv("APP_KEY")
	if !exists {
		return errors.New("APP_KEY is not set")
	}

	appName, exists := os.LookupEnv("APP_NAME")
	if !exists {
		return errors.New("APP_NAME is not set")
	}

	env, exists := os.LookupEnv("ENV")
	if !exists {
		return errors.New("ENV is not set")
	}

	asynqmonService, exists := os.LookupEnv("ASYNQMON_SERVICE")
	if !exists {
		return errors.New("ENV is not set")
	}

	appHost, exists := os.LookupEnv("APP_HOST")
	if !exists {
		return errors.New("APP_HOST is not set in .env")
	}

	AppConfig = App{
		AppName:         appName,
		AppKey:          appKey,
		Env:             env,
		AsynqmonService: asynqmonService,
		AppHost:         appHost,
	}

	return nil
}
