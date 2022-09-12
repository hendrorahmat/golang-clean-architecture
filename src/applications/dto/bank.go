package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type IBankDto interface {
	Validate() error
}

type BankList struct {
	Page    int
	PerPage int
	Keyword    string
}

func (b *BankList) Validate() error {
	if err := validation.ValidateStruct(
		b,
		validation.Field(&b.Page),
		validation.Field(&b.PerPage),
		validation.Field(&b.Keyword),
	); err != nil {
		//retErr := infra_error.NewError(infra_error.INVALID_REQUEST_RETRIEVE_PROVINCE, err)
		//retErr.SetValidationMessage(err)

		return err
	}
	return nil
}

func NewBankList(page int, perPage int, keyword string) IBankDto {
	return &BankList{
		Page: page,
		PerPage: perPage,
		Keyword: keyword,
	}
}