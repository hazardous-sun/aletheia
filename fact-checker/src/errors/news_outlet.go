package server_errors

const (
	NewsOutletNotFound          = "news outlet not found inside the database"
	NewsOutletAlreadyExists     = "news outlet already exists inside the database"
	NewsOutletTableMissing      = "news outlet table does not exist inside the database"
	NewsOutletParsingError      = "news outlet could not be parsed from the database"
	NewsOutletClosingTableError = "news outlet table could not be closed properly"
	NewsOutletNotAdded          = "news outlet was not properly added to the database"
)
