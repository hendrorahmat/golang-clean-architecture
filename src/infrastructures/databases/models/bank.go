package models

import "time"

type Bank struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
