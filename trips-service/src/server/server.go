// Package server all code related to the server
package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"trips-service.com/src/errors"
)

func Init(router http.Handler) (*http.Server, context.CancelFunc, error) {
	ctx, cancelCtx := context.WithCancel(context.Background())

	port := os.Getenv("PORT")
	if port == "" {
		return nil, cancelCtx, errors.NewEnvError("PORT")
	}

	if _, err := strconv.Atoi(port); err != nil {
		return nil, cancelCtx, errors.NewEnvError("PORT")
	}

	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		return nil, cancelCtx, errors.NewEnvError("SERVER_NAME")
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,

		BaseContext: func(l net.Listener) context.Context {
			return context.WithValue(ctx, serverName, l.Addr())
		},
	}

	return srv, cancelCtx, nil
}
