package logger

type ILogger interface {
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Debug(args ...any)
	Fatalf(format string, args ...interface{})
}

type LoggerType string

const (
	Logrus LoggerType = "logrus"
	Zap    LoggerType = "zap"
	Std    LoggerType = "std"
)

func NewLogger(t LoggerType) ILogger {
	switch t {
	case Logrus:
		return newLogrusLogger()
	default:
		return newLogrusLogger()
	}
}
