package config

import (
	"log"
	"os"
)

func LevelLogs() string {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	return logLevel
}

func SetLogLevel(level string) {

	switch level {
	case "debug":
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags | log.Llongfile)
	case "info":
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags)
	case "warn":
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	case "error":
		log.SetOutput(os.Stderr)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	default:
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags)
	}
}
