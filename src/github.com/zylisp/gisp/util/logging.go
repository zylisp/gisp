package util

import (
	"github.com/op/go-logging"
  "github.com/zylisp/gisp/common"
  "os"
)

func SetupLogging(stringLogLevel string) *logging.Logger {
	format := logging.MustStringFormatter(
		`%{color}%{time:2006-01-02T15:04:05.999Z-07:00} %{level} %{shortpkg}/%{shortfile}, %{shortfunc} %{id:03x} ▶ %{color:reset}%{message}`,)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)
	logLevel, err := logging.LogLevel(stringLogLevel)
	if err != nil {
		panic(common.LogLevelUnsupportedError)
	}
	backendLeveled.SetLevel(logLevel, "")
	logging.SetBackend(backendLeveled)
	log := GetLogger()
	log.Info("Set up logging")
	return log
}

func GetLogger() *logging.Logger {
	return logging.MustGetLogger(common.ApplicationName)
}
