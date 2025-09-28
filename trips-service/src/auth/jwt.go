// Package auth
package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Sub     string `json:"sub"`
	Purpose string `json:"purpose"`
	Minutes string `json:"minutes"`
	jwt.RegisteredClaims
}

func DecodeToken(tokenStr string, jwtSecretKey string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected alg: %v", t.Header["alg"])
		}

		return jwtSecretKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, fmt.Errorf("token is invalid: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims.Purpose != "sign_in" {
		return nil, errors.New("invalid purpose")
	}

	return claims, nil
}
