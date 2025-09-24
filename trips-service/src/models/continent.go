package models

import (
	"time"

	"gorm.io/gorm"
)

type Continent struct {
	ID        uint `gorm:"primarykey;autoIncrement"`
	code      string
	name      string
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
