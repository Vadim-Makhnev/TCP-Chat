package logger

import "github.com/sirupsen/logrus"

type logrusLogger struct {
	logger *logrus.Logger
}

func newLogrusLogger() ILogger {
	return &logrusLogger{
		logger: logrus.New(),
	}
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logrusLogger) Fatalf(fromat string, args ...interface{}) {
	l.logger.Fatalf(fromat, args...)
}
