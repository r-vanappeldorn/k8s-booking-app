// Package server all code related to the server
package server

import (
	"context"
	"net"
	"net/http"

	"trips-service.com/src/config"
)

func Init(router http.Handler, env *config.Env) (*http.Server, context.CancelFunc, error) {
	ctx, cancelCtx := context.WithCancel(context.Background())

	srv := &http.Server{
		Addr:    ":80",
		Handler: router,

		BaseContext: func(l net.Listener) context.Context {
			return context.WithValue(ctx, env.ServerName, l.Addr())
		},
	}

	return srv, cancelCtx, nil
}
