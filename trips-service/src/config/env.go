// Package config:
package config

import (
	"os"

	"trips-service.com/src/errors"
)

type Env struct {
	DBUser     string
	DBPassword string
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

	return &Env{
		DBUser:     dbUser,
		DBPassword: dbPassword,
	}, nil
}
