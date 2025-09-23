// Package controllers:
package controllers

import (
	"encoding/json"
	"net/http"

	"trips-service.com/src/router"
)

type HealthController struct {
	r *router.Router
}

func NewHealthController(r *router.Router) *HealthController {
	return &HealthController{r}
}

func (c *HealthController) Mount(r *router.Router) {
	r.Get("/health", c.Health)
}

func (c *HealthController) Health(w http.ResponseWriter, r *http.Request, ctx *router.Conext) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
