// Package testutils:
package testutils

import (
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"trips-service.com/src/config"
	"trips-service.com/src/server"
)

type TestContext struct {
	GormDB     *gorm.DB
	Mock       sqlmock.Sqlmock
	CloseSQLDB func() error
}

func InitTestServer(t *testing.T) (*http.Server, *TestContext) {
	env := &config.Env{}

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	gormDB, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      sqlDB,
			SkipInitializeWithVersion: true,
		},
	), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	srv, _, err := server.Init(env, gormDB)
	if err != nil {
		t.Fatal(err)
	}

	return srv, &TestContext{
		gormDB,
		mock,
		sqlDB.Close,
	}
}
