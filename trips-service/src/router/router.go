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

type (
	HandlerFunc func(http.ResponseWriter, *http.Request, *Conext)
	Middleware  func(HandlerFunc) HandlerFunc
)

func (r *Router) HandleWith(path string, handler HandlerFunc, method string, mws ...Middleware) {
	h := chain(handler, mws...)
	r.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		r.ctx.Logger.Info("route hit", "method", method, "path", req.URL.Path)
		if req.Method != method {
			r.ctx.Logger.Error("route not found", "method", method, "path", req.URL.Path)
			http.Error(w, "Invalid method used on route", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		h(w, req, r.ctx)
	})
}

func chain(h HandlerFunc, mws ...Middleware) HandlerFunc {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}

	return h
}

func (r *Router) Get(patern string, handler HandlerFunc, mws ...Middleware) {
	r.HandleWith(patern, handler, http.MethodGet, mws...)
}

func (r *Router) Post(patern string, handler HandlerFunc, mws ...Middleware) {
	r.HandleWith(patern, handler, http.MethodPost, mws...)
}

func (r *Router) Put(patern string, handler HandlerFunc, mws ...Middleware) {
	r.HandleWith(patern, handler, http.MethodPut, mws...)
}

func (r *Router) Patch(patern string, handler HandlerFunc, mws ...Middleware) {
	r.HandleWith(patern, handler, http.MethodPatch, mws...)
}

func (r *Router) Delete(patern string, handler HandlerFunc, mws ...Middleware) {
	r.HandleWith(patern, handler, http.MethodDelete, mws...)
}
