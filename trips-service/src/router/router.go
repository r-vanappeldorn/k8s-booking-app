// Package router
package router

import (
	"database/sql"
	"log/slog"
	"net/http"

	"trips-service.com/src/config"
)

type Conext struct {
	env    *config.Env
	logger *slog.Logger
	conn   *sql.DB
}

type Router struct {
	*PrefixMux
	ctx    *Conext
}

type HandlerFunc func(http.ResponseWriter, *http.Request, *Conext)

func (r *Router) Handle(patern string, handler HandlerFunc, method string) {
	r.HandleFunc(patern, func(w http.ResponseWriter, req *http.Request) {
		r.ctx.logger.Info("route hit", "method", method, "path", req.URL.Path)
		if req.Method != method {
			r.ctx.logger.Error("route not found", "method", method, "path", req.URL.Path)
			http.Error(w, "Invalid method used on route", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		handler(w, req, r.ctx)
	})
}

func (r *Router) Get(patern string, handler HandlerFunc) {
	r.Handle(patern, handler, http.MethodGet)
}

func (r *Router) Post(patern string, handler HandlerFunc) {
	r.Handle(patern, handler, http.MethodPost)
}

func (r *Router) Put(patern string, handler HandlerFunc) {
	r.Handle(patern, handler, http.MethodPut)
}

func (r *Router) Patch(patern string, handler HandlerFunc) {
	r.Handle(patern, handler, http.MethodPatch)
}

func (r *Router) Delete(patern string, handler HandlerFunc) {
	r.Handle(patern, handler, http.MethodDelete)
}
