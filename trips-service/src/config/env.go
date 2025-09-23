// Package config:
package config

import (
	"os"

	"trips-service.com/src/errors"
)

type Env struct {
	ServerName string
}

func InitEnv() (*Env, error) {
	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		return nil, errors.NewEnvError("SERVER_NAME")
	}

	return &Env{
		ServerName: serverName,
	}, nil
}
