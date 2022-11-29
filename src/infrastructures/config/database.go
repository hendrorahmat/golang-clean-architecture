package config

import (
	"fmt"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/utils"
	"os"
)

type BasicDBConf struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
}

type MYSQLConf struct {
	Charset   string
	ParseTime bool
	Timezone  string
}

type PostgresConf struct {
	Schema string
}

type DBDriver uint

type DatabaseConfig struct {
	ConnectionName       string
	SkipCreateConnection bool
	Driver               DBDriver
	BasicDBConf
	SSLMode string
	PostgresConf
	MYSQLConf
	MaxOpenConn            int
	MaxIdleConn            int
	MaxIdleTimeConnSeconds int64
	MaxLifeTimeConnSeconds int64
}

var DBConfig Databases

func MakeDatabaseConfig() {
	config := make(Databases)
	config = Databases{
		constants.DefaultConnectionDB: {
			Driver:               constants.POSTGRES,
			SkipCreateConnection: false,
			BasicDBConf: BasicDBConf{
				Host:     utils.GetEnv("DB_HOST"),
				Username: utils.GetEnv("DB_USERNAME"),
				Password: utils.GetEnv("DB_PASSWORD"),
				Name:     utils.GetEnv("DB_NAME"),
				Port:     utils.GetEnv("DB_PORT"),
			},
			SSLMode: "",
			PostgresConf: PostgresConf{
				Schema: os.Getenv("DB_SCHEMA"),
			},
			MaxOpenConn:            0,
			MaxIdleConn:            0,
			MaxIdleTimeConnSeconds: 0,
			MaxLifeTimeConnSeconds: 0,
		},
		"mysql": {
			SkipCreateConnection: true,
			Driver:               constants.MYSQL,
			BasicDBConf: BasicDBConf{
				Host:     utils.GetEnv("DB_HOST_2"),
				Username: utils.GetEnv("DB_USERNAME_2"),
				Password: utils.GetEnv("DB_PASSWORD_2"),
				Name:     utils.GetEnv("DB_NAME_2"),
				Port:     utils.GetEnv("DB_PORT_2"),
			},
			MYSQLConf: MYSQLConf{
				Charset:   utils.GetEnvWithDefaultValue("DB_CHARSET_2", "utf8"),
				ParseTime: true,
				Timezone:  utils.GetEnvWithDefaultValue("DB_TIMEZONE_2", "Local"),
			},
			SSLMode:                "",
			MaxOpenConn:            0,
			MaxIdleConn:            0,
			MaxIdleTimeConnSeconds: 0,
			MaxLifeTimeConnSeconds: 0,
		},
	}
	fmt.Println(config)
	DBConfig = config
}

type Databases map[ConnectionDBName]DatabaseConfig

type ConnectionDBName string
