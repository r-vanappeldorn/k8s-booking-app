package testutils

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SetAuthHeader(t *testing.T, req *http.Request) {
	t.Helper()

	secret := os.Getenv("TEST_JWT_SECRET")

	const minutes = 60
	claims := jwt.MapClaims{
		"sub":     "1",
		"purpose": "signed_in",
		"minutes": minutes,
		"exp":     time.Now().UTC().Add(time.Duration(minutes) * time.Minute).Unix(),
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tok.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("unable to create test jwt token: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
}
