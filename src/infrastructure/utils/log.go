package utils

import (
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

const Default = "default"

type DefaultFieldHook struct {
	fields map[string]interface{}
}

func (h *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *DefaultFieldHook) Fire(e *logrus.Entry) error {
	for i, v := range h.fields {
		e.Data[i] = v
	}
	return nil
}

type LogConfig struct {
	IsProduction bool
	Environment  string
	LogFileName  string
	Fields       map[string]interface{}
}

type LogOption func(*LogConfig)

func IsProduction(isProd bool) LogOption {
	return func(o *LogConfig) {
		o.IsProduction = isProd
	}
}

func LogEnvironment(env string) LogOption {
	return func(logConfig *LogConfig) {
		logConfig.Environment = env
	}
}

func LogName(logname string) LogOption {
	return func(o *LogConfig) {
		o.LogFileName = logname
	}
}

func LogAdditionalFields(fields map[string]interface{}) LogOption {
	return func(o *LogConfig) {
		o.Fields = fields
	}
}

var logOnce sync.Once
var logger *logrus.Logger

// NewLogInstance ...
func NewLogInstance(logOptions ...LogOption) *logrus.Logger {
	logOnce.Do(func() {
		rootPath := GetRootPath()
		var level logrus.Level
		logger = logrus.New()

		logConfig := &LogConfig{}
		logConfig.LogFileName = Default

		for _, logOption := range logOptions {
			logOption(logConfig)
		}

		//if it is production will output warn and error level
		if logConfig.IsProduction {
			level = logrus.WarnLevel
		} else {
			level = logrus.TraceLevel
		}

		logger.SetLevel(level)
		logger.SetOutput(colorable.NewColorableStdout())
		if logConfig.IsProduction {
			logger.SetFormatter(&logrus.JSONFormatter{
				TimestampFormat: time.RFC3339,
				PrettyPrint:     true,
				CallerPrettyfier: func(f *runtime.Frame) (string, string) {
					s := strings.Split(f.Function, ".")
					funcName := s[len(s)-1]
					_, filename := path.Split(f.File)
					return funcName, filename
				},
			})
		} else {
			logger.SetFormatter(&logrus.TextFormatter{
				TimestampFormat: time.RFC3339,
				CallerPrettyfier: func(f *runtime.Frame) (string, string) {
					s := strings.Split(f.Function, ".")
					funcname := s[len(s)-1]
					_, filename := path.Split(f.File)
					return funcname, filename
				},
			})
		}

		if !logConfig.IsProduction {
			dt := time.Now().UTC()
			rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
				Filename:   rootPath + "/storage/logs/" + dt.Format("20060102") + "-log-" + logConfig.Environment + "-" + logConfig.LogFileName,
				MaxSize:    50, // megabytes
				MaxBackups: 3,
				MaxAge:     28, //days
				Level:      level,
				Formatter: &logrus.JSONFormatter{
					TimestampFormat: time.RFC3339,
					CallerPrettyfier: func(f *runtime.Frame) (string, string) {
						s := strings.Split(f.Function, ".")
						funcname := s[len(s)-1]
						_, filename := path.Split(f.File)
						return funcname, filename
					},
				},
			})

			if err != nil {
				logger.Fatalf("Failed to initialize file rotate hook: %v", err)
			}

			logger.AddHook(rotateFileHook)
		}
		logger.AddHook(&DefaultFieldHook{logConfig.Fields})
	})
	return logger
}
