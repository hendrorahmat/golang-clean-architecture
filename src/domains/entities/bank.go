package entities

import "time"

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
