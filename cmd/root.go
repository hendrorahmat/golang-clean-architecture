package cmd

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/utils"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{}
var err error

func init() {
	if err = godotenv.Load(); err != nil {
		logrus.Error(".env not loaded, using default environment variables ", err.Error())
	}

	conf := config.Make()
	rootCommand.Use = utils.ToKebabCase(conf.App.Name)
	rootCommand.AddCommand(httpCommand)
}

func Execute() error {
	if err != nil {
		logrus.Error(".env not loaded, using default environment variables ", err.Error())
		return err
	}
	return rootCommand.Execute()
}
