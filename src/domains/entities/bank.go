package entities

type IBankListDto interface {
	Validate() error
}
type BankList struct {
	page int64
	perPage int32
}

func (b *BankList) Validate() error {
	panic("implement me")
}

func NewBankListDTO() IBankListDto {
	return &BankList{}
}
