package router

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"trips-service.com/src/config"
)

func Init(env *config.Env, gormDB *gorm.DB) *Router {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Error("route not found", "method", r.Method, "path", r.URL.Path)

		res := map[string]string{
			"error": "route not found",
		}

		json.NewEncoder(w).Encode(res)
	})

	prefixMux := NewPrefixMux(
		"/api/trips",
		mux,
	)

	ctx := &Conext{
		env,
		logger,
		gormDB,
		validator.New(),
	}

	return &Router{
		prefixMux,
		ctx,
	}
}
