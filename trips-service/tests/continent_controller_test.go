package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	testutils "trips-service.com/test_utils"
)

func TestCreateContinentShouldGet401(t *testing.T) {
	godotenv.Load(".env.test")
	body := bytes.NewBufferString(`{"code":"NL","name":"The Netherlands"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/trips/continent", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	srv, ctx := testutils.InitTestServer(t)
	defer ctx.CloseSQLDB()

	srv.Handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCreateContinent(t *testing.T) {
	godotenv.Load(".env.test")
	body := bytes.NewBufferString(`{"code":"NL","name":"The Netherlands"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/trips/continent", body)
	req.Header.Set("Content-Type", "application/json")
	testutils.SetAuthHeader(t, req)

	w := httptest.NewRecorder()

	srv, ctx := testutils.InitTestServer(t)

	defer ctx.CloseSQLDB()

	ctx.Mock.ExpectBegin()
	ctx.Mock.ExpectExec(`INSERT INTO .*continent.*\(.*code.*,.*name.*,.*deleted_at.*\) VALUES \(\?,\?\,\?\)`).
		WithArgs("NL", "The Netherlands", nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	ctx.Mock.ExpectCommit()
	srv.Handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateContinentValidation(t *testing.T) {
	godotenv.Load(".env.test")
	body := bytes.NewBufferString(`{"code":"N","name":"NL"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/trips/continent", body)
	req.Header.Set("Content-Type", "application/json")
	testutils.SetAuthHeader(t, req)

	w := httptest.NewRecorder()

	srv, ctx := testutils.InitTestServer(t)

	defer ctx.CloseSQLDB()

	srv.Handler.ServeHTTP(w, req)

	var res []map[string]string
	json.Unmarshal(w.Body.Bytes(), &res)

	if len(res) < 1 {
		t.Fatal("error response was empty")
	}

	messageByField := make(map[string]string)
	for _, e := range res {
		messageByField[e["field"]] = e["message"]
	}

	assert.Equal(t, messageByField["Code"], "Code must exactly be 2 characters long")
	assert.Equal(t, messageByField["Name"], "Name must be atleast 4 characters long")
}

func TestCreateContinentAlreadyExistsInDB(t *testing.T) {
	godotenv.Load(".env.test")
	body := bytes.NewBufferString(`{"code":"NL","name":"The Netherlands"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/trips/continent", body)
	req.Header.Set("Content-Type", "application/json")
	testutils.SetAuthHeader(t, req)

	w := httptest.NewRecorder()

	srv, ctx := testutils.InitTestServer(t)
	defer ctx.CloseSQLDB()

	ctx.Mock.ExpectBegin()
	ctx.Mock.ExpectExec(`INSERT INTO .*continent.*\(.*code.*,.*name.*,.*deleted_at.*\) VALUES \(\?,\?\,\?\)`).
		WithArgs("NL", "The Netherlands", nil).
		WillReturnError(&mysql.MySQLError{Number: 1062, Message: "Duplicate entry 'NL' for key 'continent.code'"})
	ctx.Mock.ExpectRollback()

	srv.Handler.ServeHTTP(w, req)

	fmt.Printf("%+v\n", w.Body)

	var res []map[string]string
	json.Unmarshal(w.Body.Bytes(), &res)

	if len(res) < 1 {
		t.Fatal("error response was empty")
	}

	assert.Equal(t, "code", res[0]["field"])
	assert.Equal(t, "Country already exists", res[0]["message"])
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
