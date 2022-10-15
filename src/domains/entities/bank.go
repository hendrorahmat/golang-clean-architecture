package entities

import (
	"encoding/json"
	"time"
)

type IBankListDto interface {
	Validate() error
}
type Bank struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Bank) Validate() error {
	panic("implement me")
}

func NewBankListDTO() IBankListDto {
	return &Bank{}
}

func MakeBankEntity(
	name string,
	createdAt time.Time,
	updatedAt time.Time,
) (*Bank, error) {
	return &Bank{
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (b *Bank) MarshalJSON() ([]byte, error) {
	respon, err := json.Marshal(struct {
		Name      string
		CreatedAt string
		UpdatedAt string
	}{
		Name:      b.Name,
		CreatedAt: b.CreatedAt.String(),
		UpdatedAt: b.UpdatedAt.String(),
	})
	if err != nil {
		return nil, err
	}
	return respon, nil
}
