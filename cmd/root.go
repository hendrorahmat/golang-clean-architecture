package cmd

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/utils"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{}
var err error

func Execute() error {
	utils.LoadEnv()
	conf := config.Make()
	rootCommand.Use = utils.ToKebabCase(conf.App.Name)
	rootCommand.AddCommand(httpCommand)
	rootCommand.AddCommand(rsaGenerator)
	return rootCommand.Execute()
}
