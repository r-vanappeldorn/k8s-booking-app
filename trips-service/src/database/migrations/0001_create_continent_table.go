package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"trips-service.com/src/models"
)

func migratiom0001CreateContinentTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "0001_create_continent_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Continent{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("continent")
		},
	}
}
