package repositories_mysql

import (
	"context"
	"fmt"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/repositories"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/repositories/gorm_types"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases/models"
)

type GormBankRepository struct {
	gorm_types.TransactionRepository
}

func (g *GormBankRepository) GetBankList(ctx context.Context, filter *repositories.BankRepositoryFilter) {
	//data := g.FindAll(ctx, models.Bank{})
	//data.Error()
	//data.Error()
	var bankList []models.Bank
	//ctxTest := context.Background()
	g.FindAll(ctx, &bankList)
	fmt.Println("From Mysql")
	fmt.Println(*&bankList[0].Name)
}