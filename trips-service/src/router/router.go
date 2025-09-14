// Package router
package router

import (
	"net/http"

	"trips-service.com/src/config"
	"trips-service.com/src/server"
)

type Router struct {
	*server.PrefixMux
	Env *config.Env
}

type HandlerFunc func(http.ResponseWriter, *http.Request, *config.Env)

func (r *Router) Handle(patern string, handler HandlerFunc, method string) {
	r.HandleFunc(patern, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			http.Error(w, "Invalid method used on route", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		handler(w, req, r.Env)
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
