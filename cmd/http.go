package cmd

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/src/bootstrap"
	"github.com/hendrorahmat/golang-clean-architecture/src/domains/repositories"
	"github.com/spf13/cobra"
)

var httpCommand = &cobra.Command{
	Use:   "http",
	Short: "Run HTTP Api",
	RunE: func(cmd *cobra.Command, args []string) error {
		application := bootstrap.Boot()
		ctx := context.Background()
		application.GetRepositoryCustomConnection("mysql").BankRepository.GetBankList(ctx, &repositories.BankRepositoryFilter{})
		//application.GetRepository().BankRepository.GetBankList(ctx, &repositories.BankRepositoryFilter{})
		//repo := databases.(db.Postgres.DB(), logger)
		//repo.BankRepository.GetBankList(ctx, &repositories.BankRepositoryFilter{})
		//repo.GetBankList(ctx, &repositories.BankRepositoryFilter{})

		return nil
	},
}
