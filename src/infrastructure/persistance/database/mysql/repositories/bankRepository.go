package repositories_mysql

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/entities"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/persistance/database/models"
)

type MysqlBankRepository struct {
	repositories.ITransactionRepository
}

func (g *MysqlBankRepository) GetBankList(ctx context.Context, filter *repositories.BankRepositoryFilter) ([]entities.Bank, error) {
	var bankEntities []entities.Bank

	var bankList []models.Bank

	g.FindAll(ctx, &bankList)

	for _, bank := range bankList {
		bankEntity, _ := bank.ToEntity()
		bankEntities = append(bankEntities, *bankEntity)
	}
	return bankEntities, nil
}
