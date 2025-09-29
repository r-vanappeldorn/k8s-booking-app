package models

import (
	"time"

	"gorm.io/gorm"
)

type Continent struct {
	ID        uint           `json:"id" gorm:"primarykey;autoIncrement"`
	Code      string         `json:"code" gorm:"type:char(3);not null;uniqueIndex"`
	Name      string         `json:"name" gorm:"type:varchar(200);not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (c *Continent) TableName() string {
	return "continent"
}
