package repositories_postgres

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/entities"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/repositories"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/repositories/gorm_types"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases/models"
)

type GormBankRepository struct {
	gorm_types.TransactionRepository
}

func (g *GormBankRepository) GetBankList(ctx context.Context, filter *repositories.BankRepositoryFilter) ([]entities.Bank, error) {
	var bankEntities = make([]entities.Bank, 0)

	var bankList []models.Bank

	g.FindAll(ctx, &bankList)

	for _, bank := range bankList {
		bankEntity, _ := bank.ToEntity()
		bankEntities = append(bankEntities, *bankEntity)
	}
	return bankEntities, nil
}
