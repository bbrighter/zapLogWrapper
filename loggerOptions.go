package logger

import "os"

type LoggerOptions struct {
	level     LogLevel
	folder    string
	fileName  string
	namespace string
	service   string
}

func NewLoggerOptions() *LoggerOptions {
	homeDir, _ := os.UserHomeDir()
	return &LoggerOptions{
		level:     Info,
		folder:    homeDir,
		fileName:  "logs.log",
		namespace: "namespace",
	}
}

func (o *LoggerOptions) SetLevel(level LogLevel) {
	o.level = level
}

func (o *LoggerOptions) SetFolder(folder string) {
	o.folder = folder
}

func (o *LoggerOptions) SetFileName(fileName string) {
	o.fileName = fileName
}

type ServiceType string

const (
	UI  ServiceType = "ui"
	API ServiceType = "api"
)

func (o *LoggerOptions) SetServiceAndAppName(namespace string, serviceType ServiceType) {
	o.namespace = namespace
	o.service = namespace + "-" + string(serviceType)
}
