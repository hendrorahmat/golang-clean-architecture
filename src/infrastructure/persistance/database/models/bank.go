package models

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	"time"
)

type Bank struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (b Bank) ToEntity() (*entities.Bank, error) {
	entity, err := entities.MakeBankEntity(b.Name, b.CreatedAt, b.UpdatedAt)
	if err != nil {
		return &entities.Bank{}, err
	}
	return entity, nil
}
