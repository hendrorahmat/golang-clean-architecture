package usecases

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/repositories"
)

type IBankUsecase interface {
	GetListBank(ctx context.Context)
}

type bankUsecase struct {
	repository repositories.IBankRepository
}

func (b bankUsecase) GetListBank(ctx context.Context) {
	b.repository.GetBankList(ctx, nil)
}

func NewBankUsecase(repo repositories.IBankRepository) IBankUsecase {
	return bankUsecase{
		repository: repo,
	}
}
