package repositories

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/entities"
)

type BankRepositoryFilter struct {
	Keyword string
	page    int64
	perPage uint16
}

type IBankRepository interface {
	GetBankList(ctx context.Context, filter *BankRepositoryFilter) ([]entities.Bank, error)
}
