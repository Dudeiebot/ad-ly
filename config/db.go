package config

import (
	"errors"
	"fmt"
	"os"
)

var dbConfig DB

type DB struct {
	DBName      string
	DBPassword  string
	DBUsername  string
	DBPort      string
	DBHost      string
	RedisHost   string
	RedisPass   string
	RedisPort   string
	RedisUser   string
	RedisScheme string
	RedisAddr   string
}

func loadDbEnv() error {
	dBHost, exists := os.LookupEnv("DB_HOST")
	if !exists {
		return errors.New("DB_HOST not in .env")
	}

	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		return errors.New("DB_NAME not in .env")
	}

	dBPassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		return errors.New("DB_PASSWORD not in .env")
	}

	dBUsername, exists := os.LookupEnv("DB_USERNAME")
	if !exists {
		return errors.New("DB_USERNAME not in .env")
	}

	dBPort, exists := os.LookupEnv("DB_PORT")
	if !exists {
		return errors.New("DB_PORT not in .env")
	}

	redisHost, exists := os.LookupEnv("REDIS_HOST")
	if !exists {
		return errors.New("REDIS_HOST not in .env")
	}

	redisPass, exists := os.LookupEnv("REDIS_PASS")
	if !exists {
		return errors.New("REDIS_PASS not in .env")
	}

	redisPort, exists := os.LookupEnv("REDIS_PORT")
	if !exists {
		return errors.New("REDIS_PORT not in .env")
	}

	redisUser, exists := os.LookupEnv("REDIS_USER")
	if !exists {
		return errors.New("REDIS_USER not in .env")
	}

	redisScheme, exists := os.LookupEnv("REDIS_SCHEME")
	if !exists {
		return errors.New("REDIS_SCHEME not in .env")
	}

	dbConfig = DB{
		DBName:      dbName,
		DBPassword:  dBPassword,
		DBUsername:  dBUsername,
		DBPort:      dBPort,
		DBHost:      dBHost,
		RedisHost:   redisHost,
		RedisPass:   redisPass,
		RedisPort:   redisPort,
		RedisUser:   redisUser,
		RedisScheme: redisScheme,
		RedisAddr:   fmt.Sprintf("%s:%s", redisHost, redisPort),
	}

	return nil
}
