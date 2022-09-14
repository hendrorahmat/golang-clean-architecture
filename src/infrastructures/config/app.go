package config

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/constants"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"sync"
)

type AppConf struct {
	Environment             string
	Name                    string
	Key                     string
	GracefulShutdownTimeout int
}

type HttpConf struct {
	Port       string
	XRequestID string
	Timeout    int
}

type MongoDbConf struct {
	Dsn string
}

type RedisConf struct {
	Address  string
	Password string
	Db       int
}

type LogConf struct {
	Name   string
	Logger *logrus.Logger
}

// Config ...
type Config struct {
	App      AppConf
	MongoDb  MongoDbConf
	Redis    RedisConf
	Http     HttpConf
	Log      LogConf
	Database Databases
}

var appConfigOnce sync.Once
var appConfig *Config

// Make builds a appConfig value based on .env file.
func Make() *Config {
	appConfigOnce.Do(func() {
		gracefulShutdownTimeout, err := strconv.Atoi(os.Getenv("GRACEFUL_SHUTDOWN_TIMEOUT"))
		app := AppConf{
			Environment:             strings.ToLower(os.Getenv("APP_ENV")),
			Name:                    os.Getenv("APP_NAME"),
			Key:                     os.Getenv("APP_KEY"),
			GracefulShutdownTimeout: gracefulShutdownTimeout,
		}

		mongodb := MongoDbConf{
			Dsn: os.Getenv("MONGO_DSN"),
		}

		http := HttpConf{
			Port:       os.Getenv("HTTP_PORT"),
			XRequestID: os.Getenv("HTTP_REQUEST_ID"),
		}

		log := LogConf{
			Name:   os.Getenv("LOG_NAME"),
			Logger: logrus.New(),
		}

		if app.Key == "" {
			logrus.Fatalf("Please generate random string and set to APP_KEY .env variable")
			panic("Please generate random string and set to APP_KEY")
		}

		db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		redis := RedisConf{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			Db:       db,
		}

		// set default env to local
		if app.Environment == "" {
			app.Environment = "local"
		}

		// set default port for HTTP
		if http.Port == "" {
			http.Port = string(constants.DefaultPort)
		}

		httpTimeout, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT"))
		if err == nil {
			http.Timeout = httpTimeout
		}

		if os.Getenv("DB_DRIVER") == "" {
			panic(constants.DbDriverNotFound)
		}

		appConfig = &Config{
			App:      app,
			MongoDb:  mongodb,
			Http:     http,
			Redis:    redis,
			Log:      log,
			Database: DBConfig,
		}
	})

	return appConfig
}
