package models

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	ID         uint           `gorm:"primarykey;autoIncrement"`
	Name       string         `gorm:"type:varchar(200);not null"`
	PostalCode string         `gorm:"type:varchar(200);not null"`
	Latitude   float32        `gorm:"type:decimal;not null"`
	Longitude  float32        `gorm:"type:decimal;not null"`
	CountryID  uint           `gorm:"not null;index"`
	Country    Country        `gorm:"foreignKey:CountryID;references:ID;constraint:OnUpdate:RESTRICT;OnDelete:RESTRICT"`
	CreatedAt  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP;OnUpdate:CURRENT_TIMESTAMP"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (l *Location) TableName() string {
	return "location"
}
