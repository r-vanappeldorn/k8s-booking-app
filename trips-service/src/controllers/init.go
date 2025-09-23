package controllers

import "trips-service.com/src/router"

func Init(r *router.Router) {
	NewHealthController(r).Mount(r)
}
