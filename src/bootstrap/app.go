package bootstrap

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/utils"
	"github.com/sirupsen/logrus"
	"runtime"
)

type App struct {
	databases *databases.Connections
	config    *config.Config
	logger    *logrus.Logger
}

func (a *App) GetActiveConnection() databases.IDB {
	return a.databases.Connection[constants.ActiveConnectionDb]
}

func (a *App) SetActiveConnectionDB(connectionName string) {
	a.databases.Connection[constants.ActiveConnectionDb] = a.databases.Connection[connectionName]
}

func (a *App) GetConnections() *databases.Connections {
	return a.databases
}

func (a *App) GetConnection(name string) databases.IDB {
	return a.databases.Connection[name]
}

func (a *App) GetLogger() *logrus.Logger {
	return a.logger
}

func (a *App) GetRepository() *databases.Repository {
	return databases.InjectRepository(a.GetActiveConnection().DB(), a.logger)
}

func (a *App) GetRepositoryCustomConnection(connectionName string) *databases.Repository {
	if _, ok := a.databases.Connection[connectionName]; !ok {
		a.logger.Fatalf(constants.ConnectionNotEstablished)
		panic(constants.ConnectionNotEstablished)
	}

	return databases.InjectRepository(a.GetConnection(connectionName).DB(), a.logger)
}

func (a *App) GetConfig() *config.Config {
	return a.config
}

type IApp interface {
	GetRepository() *databases.Repository
	GetRepositoryCustomConnection(connectionName string) *databases.Repository
	GetConfig() *config.Config
	GetLogger() *logrus.Logger
	GetConnections() *databases.Connections
	GetConnection(name string) databases.IDB
	SetActiveConnectionDB(connectionName string)
	GetActiveConnection() databases.IDB
}

func Boot() IApp {
	utils.LoadEnv()
	conf := config.Make()
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["total-goroutine"] = runtime.NumGoroutine()
	m["service"] = utils.ToKebabCase(conf.App.Name)

	isProduction := false

	if conf.App.Environment == constants.PRODUCTION {
		isProduction = true
	}

	logger := utils.NewLogInstance(
		utils.LogName(conf.Log.Name),
		utils.IsProduction(isProduction),
		utils.LogAdditionalFields(m),
		utils.LogEnvironment(utils.GetEnvWithDefaultValue("APP_ENV", "local")),
	)

	db := databases.MakeDatabase(conf.Database, logger)
	app := &App{
		config:    conf,
		logger:    logger,
		databases: db,
	}
	return app
}
