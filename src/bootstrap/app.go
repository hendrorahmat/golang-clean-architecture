package bootstrap

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases"
	"github.com/sirupsen/logrus"
)

type App struct {
	databases *databases.Connections
	config    *config.Config
	logger    *logrus.Logger
}

func (a *App) GetRepository() *databases.Repository {
	db := a.databases.Connection["default"].DB()
	return databases.InjectRepository(db, a.logger)
}

func (a *App) GetRepositoryCustomConnection(connectionName string) *databases.Repository {
	if _, ok := a.databases.Connection[connectionName]; !ok {
		a.logger.Fatalf(constants.ConnectionNotEstablished)
		panic(constants.ConnectionNotEstablished)
	}

	db := a.databases.Connection[connectionName].DB()
	return databases.InjectRepository(db, a.logger)
}

func (a *App) GetConfig() *config.Config {
	return a.config
}

type AppContract interface {
	GetRepository() *databases.Repository
	GetRepositoryCustomConnection(connectionName string) *databases.Repository
	GetConfig() *config.Config
}

func Boot() AppContract {
	conf := config.Make()
	logger := config.NewLogger(conf)
	db := databases.MakeDatabase(conf.Database, logger)

	return &App{
		config:    conf,
		logger:    logger,
		databases: db,
	}
}
