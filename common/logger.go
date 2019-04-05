package common

import (
	"flag"

	log "github.com/zylisp/zylog/logger"
)

func SetupLogger(level string) {
	if flag.Lookup("test.v") != nil {
		setupTestLogger()
	} else {
		setupMainLogger(level)
	}
}

func setupMainLogger(level string) {
	log.SetupLogging(&log.ZyLogOptions{
		Colored:      true,
		Level:        level,
		Output:       "stdout",
		ReportCaller: true,
	})
}

func setupTestLogger() {
	log.SetupLogging(&log.ZyLogOptions{
		Colored:      false,
		Level:        "fatal",
		Output:       "stdout",
		ReportCaller: false,
	})
}
