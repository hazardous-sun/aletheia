package repositories

import (
	"database/sql"
)

type databaseRepository struct {
    connection *sql.DB
}

func NewDatabaseRepository(connection *sql.DB) databaseRepository {
	return databaseRepository{
		connection: connection,
	}
}

