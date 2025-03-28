package custom_errors

import (
	"fmt"
	"log"
)

const (
	LanguageNotFound          = "language not found inside the database"
	LanguageAlreadyExists     = "language already exists inside the database"
	LanguageTableMissing      = "language table does not exist inside the database"
	LanguageParsingError      = "language row could not be parsed from the database"
	LanguageClosingTableError = "language table could not be closed properly"
	LanguageNotAdded          = "language was not properly added to the database"
)

const (
	NewsOutletNotFound          = "news outlet not found inside the database"
	NewsOutletAlreadyExists     = "news outlet already exists inside the database"
	NewsOutletTableMissing      = "news outlet table does not exist inside the database"
	NewsOutletParsingError      = "news outlet could not be parsed from the database"
	NewsOutletClosingTableError = "news outlet table could not be closed properly"
	NewsOutletNotAdded          = "news outlet was not properly added to the database"
)

const (
	EmptyIdError   = "id cannot be empty"
	InvalidIdError = "id should be an integer"
	EmptyNameError = "name cannot be empty"
)

const (
	InfoLevel    = "info"
	WarningLevel = "warning"
	ErrorLevel   = "error"
)

const (
	errorColor   = "\033[91m"
	warningColor = "\033[93m"
)

func CustomLog(message string, level string) {
	switch level {
	case InfoLevel:
		log.Println(fmt.Sprintf("info: %s", message))
	case WarningLevel:
		log.Println(fmt.Sprintf("%swarning: %s", warningColor, message))
	case ErrorLevel:
		log.Println(fmt.Sprintf("%serror: %s", errorColor, message))
	}
}
