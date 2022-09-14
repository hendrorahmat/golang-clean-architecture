package usecases

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/entities"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/repositories"
)

type IBankUsecase interface {
	GetListBank(ctx context.Context) ([]entities.Bank, error)
}

type BankUsecase struct {
	Repository repositories.IBankRepository
}

func (b *BankUsecase) GetListBank(ctx context.Context) ([]entities.Bank, error) {
	return b.Repository.GetBankList(ctx, nil)
}
