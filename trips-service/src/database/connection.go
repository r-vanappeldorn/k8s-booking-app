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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		env.DBUser,
		env.DBPassword,
		"trips-service-db-srv.staging-ns.svc.cluster.local",
		"3306",
		"trips",
	)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to create db connection: %w", err)
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(100)

	return sqlDB, nil
}
