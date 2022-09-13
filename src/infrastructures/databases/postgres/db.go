package postgres

import (
	"database/sql"
	"fmt"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"os"
	"strconv"
	"time"
)

type connectionPostgresDB struct {
	database *gorm.DB
	sqlDb    *sql.DB
}

type IPostgresDB interface {
	SqlDB() *sql.DB
	DB() *gorm.DB
}

func NewPostgresDB() *connectionPostgresDB {
	return &connectionPostgresDB{}
}

func (p *connectionPostgresDB) NewConnection(config config.DatabaseConfig, log *logrus.Logger) IPostgresDB {
	logger := log
	logger.Info(fmt.Sprintf("Creating Connection %s ...", config.ConnectionName))
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host,
		config.Username,
		config.Password,
		config.Name,
		config.Port,
	)

	if config.Password == "" {
		dsn = fmt.Sprintf(
			"host=%s user=%s dbname=%s port=%s sslmode=disable",
			config.Host,
			config.Username,
			config.Name,
			config.Port,
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

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: gormLog.Default.LogMode(gormLog.Warn),
	})

	logger.Info("Connecting to Postgres database...")
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
	logger.Info(fmt.Sprintf("Database %s Connected Successfully!", config.Name))
	sqlDB.Ping()
	p.sqlDb = sqlDB
	p.database = db

	return p
}

func (p *connectionPostgresDB) SqlDB() *sql.DB {
	return p.sqlDb
}

func (p *connectionPostgresDB) DB() *gorm.DB {
	return p.database
}
