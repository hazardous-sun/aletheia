package repositories

import (
	custom_errors "ai-fact-checker/server-api/errors"
	"ai-fact-checker/server-api/models"
	"database/sql"
	"errors"
)

type LanguageRepository struct {
	connection *sql.DB
}

func NewLanguageRepository(connection *sql.DB) LanguageRepository {
	return LanguageRepository{
		connection: connection,
	}
}

// Create ----------------------------------------------------------------------------------

func (lr *LanguageRepository) CreateLanguage(newsOutlet models.NewsOutlet) (int, error) {
	_, err := lr.connection.Prepare("INSERT INTO languages (name) VALUES ($1) RETURNING id")
	
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, errors.New(custom_errors.LanguageNotFound)
		}
		// TODO implement missing handler for the scenario where the language already exists
		return -1, err
	}
	
	return -1, nil
}

// Read ------------------------------------------------------------------------------------

func (lr *LanguageRepository) GetLanguages() ([]models.Language, error) {
	query := "SELECT * FROM languages"
	rows, err := lr.connection.Query(query)

	if err != nil {
		custom_errors.CustomLog(err.Error(), custom_errors.ErrorLevel)
		return []models.Language{}, errors.New(custom_errors.LanguageTableMissing)
	}

	var languageList []models.Language
	var languageObj models.Language

	for rows.Next() {
		err = rows.Scan(
			&languageObj.Id,
			&languageObj.Name,
		)

		if err != nil {
			custom_errors.CustomLog(err.Error(), custom_errors.ErrorLevel)
			return []models.Language{}, errors.New(custom_errors.LanguageParsingError)
		}

		languageList = append(languageList, languageObj)
	}

	err = rows.Close()

	if err != nil {
		custom_errors.CustomLog(err.Error(), custom_errors.ErrorLevel)
		return []models.Language{}, errors.New(custom_errors.LanguageClosingTableError)
	}

	return languageList, nil
}