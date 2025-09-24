package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"trips-service.com/src/config"
	"trips-service.com/src/server"
)

func initTestServer() (*http.Server, sqlmock.Sqlmock, func() error, error) {
	env := &config.Env{}

	sqlDB, mock, err := sqlmock.New()

	gormDB , err := gorm.Open(mysql.New(
		mysql.Config{
			Conn: sqlDB,
		},
	), &gorm.Config{})
	if err != nil {
		return nil, nil, nil, err
	}

	srv, _, err := server.Init(env, gormDB)

	return srv, mock, sqlDB.Close, err
}

func TestHealthRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/trips/health", nil)
	w := httptest.NewRecorder()

	srv, _, close, err := initTestServer()
	if err != nil {
		t.Fatal(err)
	}

	defer close()

	srv.Handler.ServeHTTP(w, req)

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
	srv, _, close, err := initTestServer()
	if err != nil {
		t.Fatal(err)
	}

	defer close()

	req := httptest.NewRequest(http.MethodPost, "/api/trips/health", nil)
	w := httptest.NewRecorder()

	srv.Handler.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404 not found, got: %d", w.Code)
	}
}
