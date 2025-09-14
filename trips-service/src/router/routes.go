package router

import (
	"encoding/json"
	"net/http"

	"trips-service.com/src/config"
	"trips-service.com/src/server"
)

func Init(env *config.Env) http.Handler {
	prefixMux := server.NewPrefixMux(
		"/api/trips",
		http.NewServeMux(),
	)

	router := &Router{prefixMux, env}

	router.Get("/health", func(w http.ResponseWriter, r *http.Request, env *config.Env) {
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	return router.Mux
}
