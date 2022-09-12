package mysql

import (
	"database/sql"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type mysqlDB struct {
	Database *gorm.DB
	SqlDb    *sql.DB
}

type IMysqlDB interface {
	SqlDB() *sql.DB
	DB() *gorm.DB
}

func (p *mysqlDB) SqlDB() *sql.DB {
	return p.SqlDb
}

func (p *mysqlDB) DB() *gorm.DB {
	return p.Database
}

func NewConnection(config config.PostgresDbConf, log *logrus.Logger) IMysqlDB {
	return &mysqlDB{}
}
