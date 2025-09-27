// Package router
package router

import (
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"trips-service.com/src/config"
)

type Conext struct {
	Env       *config.Env
	Logger    *slog.Logger
	GormDB    *gorm.DB
	Validator *validator.Validate
}

type Router struct {
	*PrefixMux
	ctx *Conext
}

type HandlerFunc func(http.ResponseWriter, *http.Request, *Conext)

func (r *Router) Handle(patern string, handler HandlerFunc, method string) {
	r.HandleFunc(patern, func(w http.ResponseWriter, req *http.Request) {
		r.ctx.Logger.Info("route hit", "method", method, "path", req.URL.Path)
		if req.Method != method {
			r.ctx.Logger.Error("route not found", "method", method, "path", req.URL.Path)
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
