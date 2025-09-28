// Package config:
package config

import (
	"os"

	"trips-service.com/src/errors"
)

type Env struct {
	DBUser       string
	DBPassword   string
	JwtSecretKey string
}

func InitEnv() (*Env, error) {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, errors.NewEnvError("DB_USER")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbUser == "" {
		return nil, errors.NewEnvError("DB_PASSWORD")
	}

	JwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if dbUser == "" {
		return nil, errors.NewEnvError("JWT_SECRET_KEY")
	}

	return &Env{
		DBUser:       dbUser,
		DBPassword:   dbPassword,
		JwtSecretKey: JwtSecretKey,
	}, nil
}
