package repositories

import (
	"context"
)

type BankRepositoryFilter struct {
	Keyword string
	page int64
	perPage uint16
}

type IBankRepository interface {
	GetBankList(ctx context.Context, filter *BankRepositoryFilter)
}