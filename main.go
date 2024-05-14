package main

import (
	"os"
	"spectacle/cmd"
	"spectacle/logger"
	"strings"
)

func init() {
	// TODO: Load the .env file and decide if colors are needed in logs
	ifColorsStr := os.Getenv("COLORED_LOGS")
	ifColorsStr = strings.ToLower(ifColorsStr)
	ifColors := false
	if ifColorsStr == "yes" || ifColorsStr == "true" {
		ifColors = true
	}

	logLevelStr := os.Getenv("LOG_LEVEL")
	logLevelStr = strings.ToLower(logLevelStr)
	logLevel := "debug"
	if logLevelStr != "" {
		logLevel = logLevelStr
	}
	// By default, the color are not present
	// and log level is debug
	logger.Log = logger.NewLogger(logLevel, ifColors)
	logger.Log.Debugf("Started logging with colors: %v and logLevel %v", ifColors, logLevel)
}

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Log.Errorf("%+v", err)
		os.Exit(1)
	}
}
