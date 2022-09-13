package mysql

import (
	"database/sql"
	"fmt"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"os"
	"strconv"
	"time"
)

type mysqlDB struct {
	database *gorm.DB
	sqlDb    *sql.DB
}

type IMysqlDB interface {
	SqlDB() *sql.DB
	DB() *gorm.DB
}

func (p *mysqlDB) SqlDB() *sql.DB {
	return p.sqlDb
}

func (p *mysqlDB) DB() *gorm.DB {
	return p.database
}
func NewMysqlDB() *mysqlDB {
	return &mysqlDB{}
}

func (m *mysqlDB) NewConnection(config config.DatabaseConfig, log *logrus.Logger) IMysqlDB {
	logger := log
	logger.Info(fmt.Sprintf("Creating connection %s ...", config.ConnectionName))
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.MYSQLConf.Charset,
		config.MYSQLConf.ParseTime,
		config.MYSQLConf.Timezone,
	)

	if config.Password == "" {
		dsn = fmt.Sprintf(
			"%s:@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
			config.Username,
			config.Host,
			config.Port,
			config.Name,
			config.MYSQLConf.Charset,
			config.MYSQLConf.ParseTime,
			config.MYSQLConf.Timezone,
		)
	}

	dBMaxOpenConn, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err == nil {
		config.MaxOpenConn = dBMaxOpenConn
	} else {
		config.MaxOpenConn = 100
	}

	dBMaxIdleConn, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err == nil {
		config.MaxIdleConn = dBMaxIdleConn
	} else {
		config.MaxIdleConn = 10
	}

	dBMaxIdleTimeConnSeconds, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_TIME_CONN_SECONDS"))
	if err == nil {
		config.MaxIdleTimeConnSeconds = int64(dBMaxIdleTimeConnSeconds)
	}

	dBMaxLifeTimeConnSeconds, err := strconv.Atoi(os.Getenv("DB_MAX_LIFE_TIME_CONN_SECONDS"))
	if err == nil {
		config.MaxLifeTimeConnSeconds = int64(dBMaxLifeTimeConnSeconds)
	} else {
		config.MaxLifeTimeConnSeconds = time.Hour.Milliseconds()
	}

	logger.Info("Connecting to Mysql database...")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 255,
	}), &gorm.Config{
		Logger: gormLog.Default.LogMode(gormLog.Warn),
	})

	if err != nil {
		logger.Fatalf("connect database err: %s", err)
		panic("Failed to connect to database!")
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifeTimeConnSeconds))
	sqlDB.SetConnMaxIdleTime(time.Duration(config.MaxIdleTimeConnSeconds))
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)

	if err != nil {
		logger.Fatalf("database err: %s", err)
		panic("database Error!")
	}
	logger.Info("Success Connect MySql database")
	sqlDB.Ping()
	m.sqlDb = sqlDB
	m.database = db

	return m
}
