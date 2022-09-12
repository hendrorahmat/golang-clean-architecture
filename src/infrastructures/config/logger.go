package config

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/utils"
	"github.com/sirupsen/logrus"
)

func NewLogger(conf *Config) *logrus.Logger {
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = utils.ToKebabCase(conf.App.Name)

	isProd := false

	if conf.App.Environment == constants.PRODUCTION {
		isProd = true
	}

	logger := utils.NewLogInstance(
		utils.LogName(conf.Log.Name),
		utils.IsProduction(isProd),
		utils.LogAdditionalFields(m))

	return logger
}
