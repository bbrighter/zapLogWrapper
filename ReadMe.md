# Content

This is a tiny wrapper around [zap](https://pkg.go.dev/go.uber.org/zap) to enable uniform logging for my usecases.

# Setup
Input for the logger is
* log level with possible values `debug`, `info`, `warning`, and `error`,
* a log file name which should be different for every service,
* a service name and 
* an app name which should be `<servicename>-<type>`, e.g. `service-api`.