package custom_errors

import (
	"fmt"
	"log"
)

const (
	LanguageNotFound = "language not found inside the database"
	LanguageAlreadyExists = "language already exists inside the database"
	NewsOutletNotFound = "news outlet not found inside the database"
	NewsOutletAlreadyExists = "news outlet already exists inside the database"
)

const (
	InfoLevel = "info"
	WarningLevel = "warning"
	ErrorLevel = "error"
)

const (
	ErrorColor = "\033[91m"
	WarningColor = "\033[93m"
)

func CustomLog(message string, level string) {
	switch level {
	case InfoLevel:
		log.Println(fmt.Sprintf("info: %s", message))
	case WarningLevel:
		log.Println(fmt.Sprintf("%swarning: %s", WarningColor, message))
	case ErrorLevel:
		log.Println(fmt.Sprintf("%serror: %s", ErrorColor, message))
	}
}