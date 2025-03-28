package repositories

import (
	custom_errors "ai-fact-checker/server-api/errors"
	"ai-fact-checker/server-api/models"
	"database/sql"
	"errors"
)

type databaseRepository struct {
    connection *sql.DB
}

func NewDatabaseRepository(connection *sql.DB) databaseRepository {
	return databaseRepository{
		connection: connection,
	}
}

