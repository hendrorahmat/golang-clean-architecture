package repositories

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/persistance/database/models"
)

type PostgresBankRepository struct {
	repositories.ITransactionRepository
}

func (g *PostgresBankRepository) GetBankList(ctx context.Context, filter *repositories.BankRepositoryFilter) ([]entities.Bank, error) {
	var bankEntities = make([]entities.Bank, 0)

	var bankList []models.Bank

	g.FindAll(ctx, &bankList)

	for _, bank := range bankList {
		bankEntity, _ := bank.ToEntity()
		bankEntities = append(bankEntities, *bankEntity)
	}
	return bankEntities, nil
}
