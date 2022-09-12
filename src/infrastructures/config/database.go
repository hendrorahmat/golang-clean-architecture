package config

import (
	"os"
)

type BasicDBConf struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
}

type PostgresDbConf struct {
	SSLMode                string
	Schema                 string
	MaxOpenConn            int
	MaxIdleConn            int
	MaxIdleTimeConnSeconds int64
	MaxLifeTimeConnSeconds int64
	BasicDBConf
}

type MysqlDbConf struct {
	BasicDBConf
}

type IDBConfig interface {
	GetPostgresConfig() PostgresDbConf
}

type DatabaseConfig struct {
	PostgresConfig PostgresDbConf
	MysqlDbConfig  MysqlDbConf
}

func MakeDatabaseConfig() DatabaseConfig {
	basicDbConf := BasicDBConf{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
	configPostgres := PostgresDbConf{
		BasicDBConf: basicDbConf,
		SSLMode:     os.Getenv("DB_SSL_MODE"),
		Schema:      os.Getenv("DB_SCHEMA"),
	}

	return DatabaseConfig{
		PostgresConfig: configPostgres,
		MysqlDbConfig:  MysqlDbConf{},
	}
}
