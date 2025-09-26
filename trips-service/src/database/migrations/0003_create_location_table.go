package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"trips-service.com/src/models"
)

func migratiom0003CreateLocationTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "0003_create_location_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Location{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("location")
		},
	}
}
