package logger

import (
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(opts *LoggerOptions) *zap.Logger {
	zapCfg := zap.NewProductionConfig()
	zapCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	filepath := createOrSelectLogFileFolder(opts.folder, opts.fileName)
	zapCfg.OutputPaths = []string{filepath}
	lvl, err := logLevelToZapLogLevel(opts.level)
	if err != nil {
		log.Fatal(err)
	}
	zapCfg.Level = zap.NewAtomicLevelAt(lvl)
	logger, err := zapCfg.Build()
	if err != nil {
		log.Fatalf("Logger could not be initialized with error %v", err)
	}
	logger = logger.With(
		zap.String("service", opts.namespace),
		zap.String("app", opts.service),
	)
	return logger
}

func createOrSelectLogFileFolder(logFolder string, fileName string) string {
	err := os.MkdirAll(logFolder, os.ModePerm)
	if err != nil {
		log.Fatalf("Could not create the directory for logs: %v", err)
	}
	filepath := filepath.Join(logFolder, fileName)
	return filepath
}
