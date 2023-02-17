package models

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"column:id,primaryKey"`
	Username  string    `json:"username" gorm:"column:username"`
	Password  string    `json:"password" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) ToEntity() (*entities.User, error) {
	entity, err := entities.MakeUserEntity(m.ID, m.Username, m.Password, m.CreatedAt, m.UpdatedAt, m.DeletedAt)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
