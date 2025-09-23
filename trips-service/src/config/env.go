// Package config:
package config

import (
	"os"

	"trips-service.com/src/errors"
)

type Env struct {
	ServerName string
	DbUser string
	DbPassword string
}

func InitEnv() (*Env, error) {
	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		return nil, errors.NewEnvError("SERVER_NAME")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, errors.NewEnvError("DB_USER")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbUser == "" {
		return nil, errors.NewEnvError("DB_PASSWORD")
	}

	return &Env{
		ServerName: serverName,
		DbUser: dbUser,
		DbPassword: dbPassword,
	}, nil
}
