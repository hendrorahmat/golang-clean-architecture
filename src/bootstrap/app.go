package bootstrap

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases"
	"github.com/sirupsen/logrus"
)

type App struct {
	repository *databases.Repository
	config     *config.Config
	logger     *logrus.Logger
}

func (a *App) GetRepository() *databases.Repository {
	return a.repository
}

func (a *App) GetConfig() *config.Config {
	return a.config
}

type AppContract interface {
	GetRepository() *databases.Repository
	GetConfig() *config.Config
}

func Boot() AppContract {
	conf := config.Make()
	logger := config.NewLogger(conf)
	conn := databases.MakeDatabase(conf.Database, logger)

	return &App{
		repository: databases.InjectRepository(conn.Postgres.DB(), logger),
		config:     conf,
		logger:     logger,
	}
}
