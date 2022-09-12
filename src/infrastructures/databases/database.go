package databases

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases/mysql"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases/postgres"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

type database struct {
	Mongo    interface{}
	Postgres postgres.IPostgresDB
	Mysql    mysql.IMysqlDB
	Redis    interface{}
}

func connectPostgres(pgConfig config.PostgresDbConf, log *logrus.Logger) postgres.IPostgresDB {
	return postgres.NewConnection(pgConfig, log)
}

func MakeDatabase(dbConfig config.DatabaseConfig, log *logrus.Logger) *database {
	driver := os.Getenv("DB_DRIVER")
	connection := database{}
	if driver == constants.DriverPostgres {
		connection.Postgres = connectPostgres(dbConfig.PostgresConfig, log)
	} else {
		panic("Driver DB not found!.")
	}
	return &connection
}

func (d *database) GetPostgresConnection() *gorm.DB {
	return d.Postgres.DB()
}

func (d *database) GetMysqlConnection() *gorm.DB {
	return d.Mysql.DB()
}
