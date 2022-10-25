package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Interface -.
type Interface interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatalf(message string, args ...interface{})
}

// Logger -.
type Logger struct {
	Logger *logrus.Logger
}

var _ Interface = (*Logger)(nil)

// New -.
func New(level string, serviceName string, file *os.File) *Logger {
	var l logrus.Level

	switch strings.ToLower(level) {
	case "error":
		l = logrus.ErrorLevel
	case "warn":
		l = logrus.WarnLevel
	case "info":
		l = logrus.InfoLevel
	case "debug":
		l = logrus.DebugLevel
	default:
		l = logrus.InfoLevel
	}

	logger := logrus.New()
	logger.SetOutput(file)
	logger.Level = l

	return &Logger{
		Logger: logger,
	}
}

// Debug -.
func (l *Logger) Debug(message string, args ...interface{}) {
	l.Logger.Debugf(message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	l.Logger.Infof(message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.Logger.Warnf(message, args...)
}

// Error -.
func (l *Logger) Error(message string, args ...interface{}) {
	l.Logger.Errorf(message, args)
}

// Fatal -.
func (l *Logger) Fatalf(message string, args ...interface{}) {
	l.Logger.Fatalf(message, args...)

	os.Exit(1)
}
