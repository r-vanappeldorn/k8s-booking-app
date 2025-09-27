// Package testutils:
package testutils

import (
	"net/http"

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

func InitTestServer() (*http.Server, *TestContext, error) {
	env := &config.Env{}

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      sqlDB,
			SkipInitializeWithVersion: true,
		},
	), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	srv, _, err := server.Init(env, gormDB)

	return srv, &TestContext{
		gormDB,
		mock,
		sqlDB.Close,
	}, err
}
