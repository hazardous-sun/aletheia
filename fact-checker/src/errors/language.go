package server_errors

const (
	LanguageNotFound          = "language not found inside the database"
	LanguageAlreadyExists     = "language already exists inside the database"
	LanguageTableMissing      = "language table does not exist inside the database"
	LanguageParsingError      = "language row could not be parsed from the database"
	LanguageClosingTableError = "language table could not be closed properly"
	LanguageNotAdded          = "language was not properly added to the database"
)
