package logger

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *Logger

type Logger struct {
	*logrus.Logger
	Files map[string]*os.File
}

func openLogFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.WithError(err).Error("Failed to open log file")
		return nil
	}
	return file
}

func setLogLevel(logger *logrus.Logger, level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.WithError(err).Error("Invalid log level")
		logger.SetLevel(logrus.InfoLevel) // Default to Info level if parsing fails
	} else {
		logger.SetLevel(lvl)
	}
}

func NewLogger(level string, enableColors bool) *Logger {
	baseLogger := logrus.New()
	setLogLevel(baseLogger, level)
	files := make(map[string]*os.File)
	files["info"] = openLogFile("spectacle.info")
	files["error"] = openLogFile("spectacle.error")
	files["debug"] = openLogFile("spectacle.debug")

	baseLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     enableColors,
		TimestampFormat: time.RFC3339,
	})
	return &Logger{baseLogger, files}
}

func (l *Logger) addContext(level string) *logrus.Entry {
	pc, file, _, ok := runtime.Caller(2)
	if !ok {
		return l.WithFields(logrus.Fields{
			"level": level,
		})
	}
	fn := runtime.FuncForPC(pc)
	entry := l.WithFields(logrus.Fields{
		"func":   fn.Name(),
		"module": filepath.Base(file),
		"level":  level,
	})
	entry.Logger.Out = l.Files[level]
	return entry
}

func (l *Logger) Info(msg string) {

	l.addContext("info").Info(msg)
}

func (l *Logger) Debug(msg string) {
	l.addContext("debug").Debug(msg)
}

func (l *Logger) Error(msg string) {
	l.addContext("error").Error(msg)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.addContext("info").Infof(format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.addContext("debug").Debugf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.addContext("error").Errorf(format, args...)
}
