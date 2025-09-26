package models

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID          uint           `gorm:"primarykey;autoIncrement"`
	Code        string         `gorm:"type:char(2);not null;uniqueIndex"`
	Name        string         `gorm:"varchar(200),not null"`
	ContinentID uint           `gorm:"not null;index"`
	Continent   Continent      `gorm:"foreignKey:ContinentID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	CreatedAt   time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (c *Country) TableName() string {
	return "country"
}
