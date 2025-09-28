// Package middleware:
package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"trips-service.com/src/auth"
	"trips-service.com/src/router"
	"trips-service.com/src/utils"
)

func NewAuthMidleware(next router.HandlerFunc) router.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c *router.Conext) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.JSONError(w, http.StatusUnauthorized, "missing auth token")

			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.DecodeToken(token, c.Env.JwtSecretKey)
		if err != nil {
			fmt.Println(err)
			utils.JSONError(w, http.StatusUnauthorized, "invalid token")

			return
		}

		context.WithValue(r.Context(), "user_id", claims.Sub)

		next(w, r, c)
	}
}
