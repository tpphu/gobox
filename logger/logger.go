package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logrus.Logger
}

type TextFormatter struct {
	logrus.TextFormatter
}

func New() *Logger {
	return &Logger{
		Logger: logrus.Logger{
			Out:          os.Stderr,
			Formatter:    new(TextFormatter),
			Hooks:        make(logrus.LevelHooks),
			Level:        logrus.InfoLevel,
			ExitFunc:     os.Exit,
			ReportCaller: false,
		},
	}
}
