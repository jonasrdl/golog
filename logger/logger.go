package logger

import (
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

type Logger struct {
	level LogLevel
}

func NewLogger(level LogLevel) *Logger {
	return &Logger{level: level}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(Debug, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.log(Info, format, args...)
}

func (l *Logger) Warning(format string, args ...interface{}) {
	l.log(Warning, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.log(Error, format, args...)
}

func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level >= l.level {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		message := fmt.Sprintf(format, args...)
		logLine := fmt.Sprintf("[%s] [%s] %s", timestamp, levelToString(level), message)
		fmt.Fprintln(os.Stdout, logLine)
	}
}

func levelToString(level LogLevel) string {
	switch level {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		return ""
	}
}
