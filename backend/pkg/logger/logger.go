package logger

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger provides structured logging functionality
type Logger struct {
	*logrus.Logger
}

// New creates a new logger instance
func New(level, format, output, timeFormat string) *Logger {
	logger := logrus.New()

	// Set log level
	switch level {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	// Set log format
	switch format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: timeFormat,
		})
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: timeFormat,
			FullTimestamp:   true,
		})
	default:
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
		})
	}

	// Set output
	switch output {
	case "stdout":
		logger.SetOutput(os.Stdout)
	case "stderr":
		logger.SetOutput(os.Stderr)
	case "file":
		file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.SetOutput(file)
		} else {
			logger.SetOutput(os.Stderr)
		}
	default:
		logger.SetOutput(os.Stdout)
	}

	return &Logger{Logger: logger}
}

// WithContext adds context information to the logger
func (l *Logger) WithContext(ctx interface{}) *logrus.Entry {
	return l.WithField("context", ctx)
}

// WithRequest adds HTTP request information to the logger
func (l *Logger) WithRequest(r interface{}) *logrus.Entry {
	return l.WithField("request", r)
}

// WithUser adds user information to the logger
func (l *Logger) WithUser(userID, email string) *logrus.Entry {
	return l.WithFields(logrus.Fields{
		"user_id": userID,
		"email":   email,
	})
}

// WithDatabase adds database information to the logger
func (l *Logger) WithDatabase(operation, table string) *logrus.Entry {
	return l.WithFields(logrus.Fields{
		"db_operation": operation,
		"db_table":     table,
	})
}

// WithError adds error information to the logger
func (l *Logger) WithError(err error) *logrus.Entry {
	return l.WithField("error", err.Error())
}

// WithFields adds multiple fields to the logger
func (l *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.WithFields(fields)
}

// SetOutput sets the logger output
func (l *Logger) SetOutput(output io.Writer) {
	l.Logger.SetOutput(output)
}

// SetLevel sets the logger level
func (l *Logger) SetLevel(level logrus.Level) {
	l.Logger.SetLevel(level)
}

// SetFormatter sets the logger formatter
func (l *Logger) SetFormatter(formatter logrus.Formatter) {
	l.Logger.SetFormatter(formatter)
}
