package entities

import "time"

type User struct {
	ID        string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func MakeUserEntity(
	id,
	username,
	password string,
	createdAt,
	updatedAt,
	deletedAt time.Time,
) (*User, error) {
	return &User{
		ID:        id,
		Username:  username,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}, nil
}
