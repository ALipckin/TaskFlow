package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"type:varchar(50);not null;default:'pending'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
