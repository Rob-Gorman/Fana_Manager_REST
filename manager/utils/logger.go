package utils

import (
	"fmt"
	"log"
	"os"
)

var InfoLog FanaLogger
var ErrLog FanaLogger

// type Logger interface {
// 	Printf(format string, args ...interface{})
// 	Falalf(format string, args ...interface{})
// }

type FanaLogger struct {
	Logger *log.Logger
}

func InitLoggers(infologger, errorlogger *log.Logger) {
	if infologger == nil {
		infologger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	}
	InfoLog = NewFanaLogger(infologger)

	if errorlogger == nil {
		errorlogger = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	}
	ErrLog = NewFanaLogger(errorlogger)
}

func NewFanaLogger(l *log.Logger) FanaLogger {
	return FanaLogger{Logger: l}
}

func (l *FanaLogger) Printf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Output(2, msg)
}

func (l *FanaLogger) Falalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Output(2, msg)
}
