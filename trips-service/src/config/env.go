// Package config:
package config

import (
	"os"
	"strconv"

	"trips-service.com/src/errors"
)

type Env struct {
	Port       int
	ServerName string
}

func InitEnv() (*Env, error) {
	envPort := os.Getenv("PORT")
	if envPort == "" {
		return nil, errors.NewEnvError("PORT")
	}

	port, err := strconv.Atoi(envPort)
	if err != nil {
		return nil, errors.NewEnvError("PORT")
	}

	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		return nil, errors.NewEnvError("SERVER_NAME")
	}

	return &Env{
		Port:       port,
		ServerName: serverName,
	}, nil
}
