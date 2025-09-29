package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func migration0005AlterContinentCodeColumn() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "0005_alter_continent_code_column",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec("ALTER TABLE continent MODIFY COLUMN code CHAR(3) NOT NULL").Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Exec("ALTER TABLE continent MODIFY COLUMN code CHAR(2) NOT NULL").Error
		},
	}
}
