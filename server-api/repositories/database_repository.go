package repositories

import (
	custom_errors "ai-fact-checker/server-api/errors"
	"ai-fact-checker/server-api/models"
	"database/sql"
	"errors"
)

type DatabaseRepository struct {
    connection *sql.DB
}

func NewDatabaseRepository(connection *sql.DB) DatabaseRepository {
	return DatabaseRepository{
		connection: connection,
	}
}

func (dr *DatabaseRepository) CreateLanguage(newsOutlet models.NewsOutlet) (int, error) {
	query, err := dr.connection.Prepare("INSERT INTO languages (name) VALUES ($1) RETURNING id")
	
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, errors.New(custom_errors.LanguageNotFound)
		}
		// TODO implement missing handler for the scenario where the language already exists
		return -1, err
	}
	
	return -1, nil
}
