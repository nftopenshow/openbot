// Package log /**
package log

import "github.com/openshow/openbot/twitter/src/config"

func init() {
	logLevel := config.Network.Logger.LogLevel
	logFileDir := config.Network.Logger.LogFileDir
	InitLog(logLevel, logFileDir, Stdout)
}
