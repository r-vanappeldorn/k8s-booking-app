// Package models:
package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TripStatus string

const (
	TripStatusDraft     TripStatus = "draft"
	TripStatusPublished TripStatus = "published"
	TripStatusArchived  TripStatus = "archived"
)

type Trip struct {
	ID             uint           `gorm:"primarykey;autoIncrement"`
	Title          string         `gorm:"type:varchar(200);not null"`
	Description    string         `gorm:"type:text"`
	StartsAt       time.Time      `gorm:"not null;index"`
	EndsAt         time.Time      `gorm:"not null;index"`
	Status         TripStatus     `gorm:"type:ENUM('draft','published','archived');default:'draft';not null;index"`
	Capacity       int            `gorm:"not null"`
	BasePriceCents int            `gorm:"not null"`
	Currency       string         `gorm:"type:char(3);not null"`
	LocationID     uint           `gorm:"not null;index"`
	Location       Location       `gorm:"foreignKey:LocationID;references:ID;constraint:OnUpdate:RESTRICT;OnDelete:RESTRICT"`
	CreatedAt      time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP;index"`
	UpdatedAt      time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP;OnUpdate:CURRENT_TIMESTAMP"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (s TripStatus) Isvalid() bool {
	switch s {
	case TripStatusDraft, TripStatusArchived, TripStatusPublished:
		return true
	}

	return false
}

func (s *TripStatus) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid type for TripStatus: %T", value)
	}

	status := TripStatus(str)
	if !status.Isvalid() {
		return fmt.Errorf("invalid string provided for type TripStatus: %s", status)
	}

	*s = status

	return nil
}
