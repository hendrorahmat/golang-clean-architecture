package databases

import (
	"database/sql"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases/mysql"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/databases/postgres"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

type IDB interface {
	SqlDB() *sql.DB
	DB() *gorm.DB
}

type Connection map[string]IDB

type Connections struct {
	Connection
}

var connections *Connections
var dbConnOnce sync.Once

func MakeDatabase(databases config.Databases, log *logrus.Logger) *Connections {
	dbConnOnce.Do(func() {
		listConnections := Connection{}
		var dbConnection IDB
		for name, databaseConf := range databases {
			if databaseConf.SkipCreateConnection {
				continue
			}

			databaseConf.ConnectionName = string(name)

			switch databaseConf.Driver {
			case constants.POSTGRES:
				pgCon := postgres.NewPostgresDB()
				dbConnection = pgCon.NewConnection(databaseConf, log)
			case constants.MYSQL:
				mysqlCon := mysql.NewMysqlDB()
				dbConnection = mysqlCon.NewConnection(databaseConf, log)
			}

			listConnections[string(name)] = dbConnection
		}

		listConnections[constants.ActiveConnectionDb] = listConnections[constants.DefaultConnectionDB]

		connections = &Connections{
			listConnections,
		}
	})

	return connections
}
func (c *Connections) GetConnection(connectionName string) *gorm.DB {
	return c.Connection[connectionName].DB()
}
