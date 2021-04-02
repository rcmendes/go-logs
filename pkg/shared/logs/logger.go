package logs

import (
	"fmt"
	"io"
	"log"
)

const LevelDebug = 1
const LevelInfo = 2
const LevelWarning = 4
const LevelError = 8

type Logger interface {
	Debug(message ...interface{})
	Info(message ...interface{})
	Warning(message ...interface{})
	Error(message ...interface{})
}

type logger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	level         int
}

func newLogger(output io.Writer, flags int, logLevel int) Logger {
	return &logger{
		debugLogger:   log.New(output, "DEBUG: ", flags),
		infoLogger:    log.New(output, "INFO: ", flags),
		warningLogger: log.New(output, "WARNING: ", flags),
		errorLogger:   log.New(output, "ERROR: ", flags),
		level:         logLevel,
	}
}

func (l *logger) Debug(message ...interface{}) {
	if l.level == LevelDebug {
		l.debugLogger.Output(2, fmt.Sprintln(message...))
	}
}

func (l *logger) Info(message ...interface{}) {
	if l.level <= LevelInfo {
		l.infoLogger.Output(2, fmt.Sprintln(message...))
	}
}

func (l *logger) Warning(message ...interface{}) {
	if l.level <= LevelWarning {
		l.warningLogger.Output(2, fmt.Sprintln(message...))
	}
}

func (l *logger) Error(message ...interface{}) {
	if l.level <= LevelError {
		l.errorLogger.Output(2, fmt.Sprintln(message...))
	}
}
