package router

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"trips-service.com/src/config"
	"trips-service.com/src/server"
)

func Init(env *config.Env, conn *sql.DB) http.Handler {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		logger.Error("route not found", "method", r.Method, "path", r.URL.Path)

		res := map[string]string {
			"error": "route not found",
		}

		json.NewEncoder(w).Encode(res)
	})

	prefixMux := server.NewPrefixMux(
		"/api/trips",
		mux,
	)

	ctx := &Conext{
		env, 
		logger,
		conn,
	}

	router := &Router{
		prefixMux,
		ctx,
	}

	router.Get("/health", func(w http.ResponseWriter, r *http.Request, _ *Conext) {
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	return router.Mux
}
