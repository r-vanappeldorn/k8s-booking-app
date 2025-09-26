package models

import (
	"time"

	"gorm.io/gorm"
)

type Continent struct {
	ID        uint           `gorm:"primarykey;autoIncrement"`
	code      string         `gorm:"type: char(2). not null"`
	name      string         `gorm:"type: varchar(200). not null"`
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (c *Continent) TableName() string {
	return "continent"
}
