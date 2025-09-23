package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"trips-service.com/src/config"
	"trips-service.com/src/router"
)

func TestHealthRoute(t *testing.T) {
	env := &config.Env{}

	conn, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	r := router.Init(env, conn)
	req := httptest.NewRequest(http.MethodGet, "/api/trips/health", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 ok, got: %d", w.Code)
	}

	var res map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatal("Invalid json responsse")
	}

	if res["status"] != "ok" {
		t.Fatalf("expected status to contain string 'ok', got: %s", res["status"])
	}
}

func TestHealthRouteMethod(t *testing.T) {
	env := &config.Env{}
	conn, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	r := router.Init(env, conn)

	req := httptest.NewRequest(http.MethodPost, "/api/trips/health", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404 not found, got: %d", w.Code)
	}
}
