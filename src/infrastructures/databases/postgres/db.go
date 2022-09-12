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

type ConnectionPostgresDB struct {
	DriverName string
	Database   *gorm.DB
	SqlDb      *sql.DB
}

type IPostgresDB interface {
	SqlDB() *sql.DB
	DB() *gorm.DB
}

func (p *ConnectionPostgresDB) SqlDB() *sql.DB {
	return p.SqlDb
}

func (p *ConnectionPostgresDB) DB() *gorm.DB {
	return p.Database
}

func NewConnection(config config.PostgresDbConf, log *logrus.Logger) IPostgresDB {
	loggrusLog := log
	loggrusLog.Info("Connecting to database...")

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

	if err != nil {
		loggrusLog.Fatalf("connect database err: %s", err)
		panic("Failed to connect to database!")
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifeTimeConnSeconds))
	sqlDB.SetConnMaxIdleTime(time.Duration(config.MaxIdleTimeConnSeconds))
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)

	if err != nil {
		loggrusLog.Fatalf("database err: %s", err)
		panic("Database Error!")
	}
	loggrusLog.Info("Success Connect Database")
	sqlDB.Ping()
	return &ConnectionPostgresDB{
		Database: db,
		SqlDb:    sqlDB,
	}
}
