package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"trips-service.com/src/config"
	"trips-service.com/src/database"
	"trips-service.com/src/database/migrations"
)

func main() {
	cmd := "up"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	env, err := config.InitEnv()
	if err != nil {
		log.Fatalf("unable to load env: %s", err.Error())
	}

	sqlDB, err := database.Init(env)
	if err != nil {
		log.Fatalf("unable to connect to db: %s", err.Error())
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		log.Fatalf("unable to create gorm connection: %s", err.Error())
	}

	switch cmd {
	case "up":
		err = migrations.Up(gormDB)
	case "down-one":
		err = migrations.DownOne(gormDB)
	case "to":
		fs := flag.NewFlagSet("to", flag.ExitOnError)
		id := fs.String("id", "", "migration ID to migrate to (inclusive)")
		_ = fs.Parse(os.Args[2:])
		if *id == "" {
			log.Fatal("usage: migrate to -id=0002_add_indexes")
		}
		err = migrations.To(gormDB, *id)
	default:
		fmt.Println("usage: migrate [up|down-one|to -id=...|steps -n=...]")
		os.Exit(2)
	}

	if err != nil {
		log.Fatalf("error in %s command: %s", cmd, err.Error())
	}

	log.Println("Migrations done")
}
