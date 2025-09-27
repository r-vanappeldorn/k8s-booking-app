package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"trips-service.com/src/models"
)

func migratiom0002CreateCountryTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "0002_create_continent_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Country{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("country")
		},
	}
}
