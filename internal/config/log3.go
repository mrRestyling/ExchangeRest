package config

import (
	"log"
	"os"
)

var (
	logFile *os.File
)

func LevelLogs() string {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	return logLevel
}

func SetLogLevel(level string) {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	switch level {
	case "debug":
		log.SetFlags(log.LstdFlags | log.Llongfile)
	case "info":
		log.SetFlags(log.LstdFlags)
	case "warn":
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	case "error":
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	default:
		log.SetFlags(log.LstdFlags)
	}

	log.SetOutput(logFile)
}

func CloseLogFile() {
	if logFile != nil {
		logFile.Close()
	}
}
