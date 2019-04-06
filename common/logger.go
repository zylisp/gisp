package common

import (
	"flag"

	log "github.com/zylisp/zylog/logger"
)

// SetupLogger is a dispatch function that calls specific setup functions
// depending upon such things as command line paramaters or whether the
// context is a test runner, etc.
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
