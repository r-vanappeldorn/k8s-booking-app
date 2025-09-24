// Package database:
package database

import (
	"database/sql"
	"fmt"
	"time"

	"trips-service.com/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func Init(env *config.Env) (*sql.DB, error) {
	sqlDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(accounts-service-db-srv.staging-ns.svc.cluster.local:3306)/trips", env.DbUser, env.DbPassword))
	if err != nil {
		return nil, fmt.Errorf("unable to create db connection: %w", err)
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(100)

	return sqlDB, nil
}
