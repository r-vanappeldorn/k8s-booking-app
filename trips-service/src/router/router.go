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

func (r *Router) Get(patern string, handler HandlerFunc) {
	r.HandleFunc(patern, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(w, "Invalid method used on route", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		handler(w, req, r.Env)
	})
}

func (r *Router) Post(patern string, handler HandlerFunc) {
	r.HandleFunc(patern, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Invalid method used on route", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		handler(w, req, r.Env)
	})
}

func (r *Router) Put(patern string, handler HandlerFunc) {
	r.HandleFunc(patern, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			http.Error(w, "Invalid method used on route", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		handler(w, req, r.Env)
	})
}
