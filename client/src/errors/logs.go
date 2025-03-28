package client_errors

import (
	"fmt"
	"log"
)

const (
	InfoLevel    = "info"
	WarningLevel = "warning"
	ErrorLevel   = "error"
)

const (
	defaultColor = "\033[0m"
	errorColor   = "\033[91m"
	warningColor = "\033[93m"
)

func Log(message string, level string) {
	switch level {
	case InfoLevel:
		log.Println(fmt.Sprintf("%sinfo: %s %s", defaultColor, message, defaultColor))
	case WarningLevel:
		log.Println(fmt.Sprintf("%swarning: %s %s", warningColor, message, defaultColor))
	case ErrorLevel:
		log.Println(fmt.Sprintf("%serror: %s %s", errorColor, message, defaultColor))
	}
}
