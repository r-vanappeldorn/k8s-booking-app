// Package server all code related to the server
package server

import (
	"context"
	"net"
	"net/http"

	"gorm.io/gorm"
	"trips-service.com/src/config"
	"trips-service.com/src/controllers"
	"trips-service.com/src/router"
)

type ServerContextName string

func Init(env *config.Env, gormDB *gorm.DB) (*http.Server, context.CancelFunc, error) {
	ctx, cancelCtx := context.WithCancel(context.Background())

	r := router.Init(env, gormDB)
	controllers.Init(r)

	var name ServerContextName = "trips-service"

	srv := &http.Server{
		Addr:    ":80",
		Handler: r.Mux,

		BaseContext: func(l net.Listener) context.Context {
			return context.WithValue(ctx, name, l.Addr())
		},
	}

	return srv, cancelCtx, nil
}
