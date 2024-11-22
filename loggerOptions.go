package logger

import (
	"fmt"
	"os"
	"path"
	"runtime/debug"
	"strings"
)

type LoggerOptions struct {
	level     LogLevel
	folder    string
	fileName  string
	namespace string
	service   string
}

// Creates default options with level info, folder current directory/logs, file name = "logs.log", namespace = "namespace"
func NewLoggerOptions() *LoggerOptions {
	opts := &LoggerOptions{
		level:    Info,
		fileName: "logs",
	}
	wd, _ := os.Getwd()
	opts.SetFolder(wd)
	service, serviceType, _ := getModuleName()
	opts.SetServiceAndAppName(service, ServiceType(serviceType))
	return opts
}

func (o *LoggerOptions) SetLevel(level LogLevel) {
	o.level = level
}

func (o *LoggerOptions) SetFolder(folder string) {
	o.folder = path.Join(folder, "logs")
}

// Sets the log file to filename.log
func (o *LoggerOptions) SetFileName(fileName string) {
	o.fileName = fileName + ".log"
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

func getModuleName() (string, string, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", "", fmt.Errorf("failed to read build info")
	}
	moduleName := path.Base(info.Main.Path)
	splits := strings.Split(moduleName, "-")
	service := splits[0]
	if len(splits) > 1 {
		serviceType := splits[1]
		if serviceType == "api" || serviceType == "ui" {
			return service, serviceType, nil
		}
	}
	return service, "", nil
}
