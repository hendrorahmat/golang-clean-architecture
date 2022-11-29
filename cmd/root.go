package cmd

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/bootstrap"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/utils"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{}
var err error

func Execute() error {
	application := bootstrap.Boot()
	conf := application.GetConfig()
	rootCommand.Use = utils.ToKebabCase(conf.App.Name)
	rootCommand.AddCommand(httpCommand)
	return rootCommand.Execute()
}
