package config

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/utils"
	"github.com/sirupsen/logrus"
	"sync"
)

var logOnce sync.Once
var logger *logrus.Logger

func NewLogger(conf *Config) *logrus.Logger {
	logOnce.Do(func() {
		m := make(map[string]interface{})
		m["env"] = conf.App.Environment
		m["service"] = utils.ToKebabCase(conf.App.Name)

		isProd := false

		if conf.App.Environment == constants.PRODUCTION {
			isProd = true
		}

		logger = utils.NewLogInstance(
			utils.LogName(conf.Log.Name),
			utils.IsProduction(isProd),
			utils.LogAdditionalFields(m))
	})

	return logger
}
