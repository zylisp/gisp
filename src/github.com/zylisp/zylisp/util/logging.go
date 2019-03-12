package util

import (
	"flag"
	"github.com/op/go-logging"
  "github.com/zylisp/zylisp/common"
  "os"
)

func setupTestLogger () logging.LeveledBackend {
	format := logging.MustStringFormatter("%{message}")
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)
	logLevel, _ := logging.LogLevel("critical")
	backendLeveled.SetLevel(logLevel, "")
	return backendLeveled
}

func setupNormalLogger (stringLogLevel string) logging.LeveledBackend {
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
	return backendLeveled
}

func SetupLogging(stringLogLevel string) *logging.Logger {
	logging.SetBackend(setupNormalLogger(stringLogLevel))
	log := GetLogger()
	log.Info("Set up logging")
	return log
}

func GetLogger() *logging.Logger {
	if flag.Lookup("test.v") != nil {
		logging.SetBackend(setupTestLogger())
	}
	return logging.MustGetLogger(common.ApplicationName)
}
