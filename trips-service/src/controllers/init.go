package controllers

import "trips-service.com/src/router"

func Init(r *router.Router) {
	NewHealthController(r).Mount(r)
	NewContinentController(r).Mount(r)
}
