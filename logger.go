package logger

import (
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(logLevel LogLevel, logFileName string, service string, app string) *zap.Logger {
	zapCfg := zap.NewProductionConfig()
	zapCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	filepath := createOrSelectLogFileFolder(logFileName)
	zapCfg.OutputPaths = []string{filepath}
	lvl, err := logLevelToZapLogLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	zapCfg.Level = zap.NewAtomicLevelAt(lvl)
	logger, err := zapCfg.Build()
	if err != nil {
		log.Fatalf("Logger could not be initializaed with error %v", err)
	}
	logger = logger.With(
		zap.String("service", service),
		zap.String("app", app),
	)
	return logger
}

func createOrSelectLogFileFolder(fileName string) string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not determine home directory: %v", err)
	}
	logfolder := filepath.Join(homedir, "logs")
	err = os.MkdirAll(logfolder, os.ModePerm)
	if err != nil {
		log.Fatalf("Could not create the directory for logs: %v", err)
	}
	filepath := filepath.Join(logfolder, fileName)
	return filepath
}
