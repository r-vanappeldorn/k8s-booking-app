// Package migrations
package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		migratiom0001CreateContinentTable(),
		migratiom0002CreateCountryTable(),
		migratiom0003CreateLocationTable(),
		migratiom0004CreateTripTable(),
		migration0005AlterContinentCodeColumn(),
	}
}

func Up(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, Migrations())
	return m.Migrate()
}

func DownOne(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, Migrations())
	return m.RollbackLast()
}

func To(db *gorm.DB, id string) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, Migrations())
	return m.MigrateTo(id)
}
