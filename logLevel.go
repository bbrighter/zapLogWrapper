package logger

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel string

const (
	Debug   LogLevel = "debug"
	Info    LogLevel = "info"
	Warning LogLevel = "warning"
	Error   LogLevel = "error"
)

func logLevelToZapLogLevel(lvl LogLevel) (zapcore.Level, error) {
	zapLevel := zapcore.InfoLevel
	var err error = nil
	switch string(lvl) {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warning":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	default:
		err = errors.New("invalid lvl in config " + string(lvl))
	}
	return zapLevel, err
}
