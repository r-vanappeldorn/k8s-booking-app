package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"trips-service.com/src/models"
)

func migratiom0004CreateTripTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "0003_create_trip_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Trip{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("trip")
		},
	}
}
