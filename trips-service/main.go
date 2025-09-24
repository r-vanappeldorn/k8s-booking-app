package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

  "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"trips-service.com/src/config"
	"trips-service.com/src/database"
	"trips-service.com/src/server"
)

func main() {
	env, err := config.InitEnv()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := database.Init(env)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	gormDB , err := gorm.Open(mysql.New(
		mysql.Config{
			Conn: sqlDB,
		},
	), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	
	srv, cancelCtx, err := server.Init(env, gormDB)
	if err != nil {
		log.Fatal(err)
	}

	signCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		errCh <- nil
	}()

	log.Println("Server started on port: 80")

	select {
	case <-signCtx.Done():
		log.Println("Shutdown server")
	case err := <-errCh:
		if err != nil {
			log.Printf("Server error: %v", err)
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cancelCtx()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Graceful shutdown failed: %v — forcing close", err)
		_ = srv.Close()
	}
}
