package models

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	ID         uint           `gorm:"primarykey;autoIncrement"`
	Name       string         `gorm:"type:varchar(200),not null"`
	PostalCode string         `gorm:"type:varchar(200),not null"`
	Latitude   float32        `gorm:"not null"`
	Longitude  float32        `gorm:"not null"`
	Country    Country        `gorm:"foreignKey:CountyID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	CreatedAt  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
