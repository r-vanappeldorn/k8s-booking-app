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
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(accounts-service-db-srv.staging-ns.svc.cluster.local:3306)", env.DbUser, env.DbPassword))
	if err != nil {
		return nil, fmt.Errorf("unable to create db connection: %w", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	return db, nil
}
