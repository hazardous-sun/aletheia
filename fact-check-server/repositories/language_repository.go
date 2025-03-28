package repositories

import (
	apiErrors "ai-fact-checker/errors"
	"ai-fact-checker/models"
	"database/sql"
	"errors"
	"strings"
)

type LanguageRepository struct {
	connection *sql.DB
}

func NewLanguageRepository(connection *sql.DB) *LanguageRepository {
	return &LanguageRepository{
		connection: connection,
	}
}

// Create --------------------------------------------------------------------------------------------------------------

// AddLanguage :
// Creates a new language inside the database based on the model received as parameter.
//
// Error: will throw LanguageTableMissing if the database is incorrectly set and the "languages" table is missing.
//
// Error: will throw LanguageParsingError if for some reason it is unable to parse the values it receives from the
// database.
//
// Error: will throw LanguageClosingTableError if it fails to close the database rows.
func (lr *LanguageRepository) AddLanguage(newsOutlet models.Language) (int, error) {
	query, err := lr.connection.Prepare("INSERT INTO languages (name) VALUES ($1) RETURNING id")

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			customErrors.CustomLog(customErrors.LanguageTableMissing, customErrors.ErrorLevel)
			return -1, errors.New(customErrors.LanguageTableMissing)
		}
		return -1, err
	}

	var id int
	name := strings.ToLower(newsOutlet.Name)
	err = query.QueryRow(name).Scan(&id)

	if err != nil {
		return -1, errors.New(customErrors.LanguageParsingError)
	}

	err = query.Close()

	if err != nil {
		customErrors.CustomLog(customErrors.LanguageClosingTableError, customErrors.ErrorLevel)
		return -1, errors.New(customErrors.LanguageClosingTableError)
	}

	return id, nil
}

// Read ----------------------------------------------------------------------------------------------------------------

// GetLanguages :
// Returns all the languages stored in the database. Even though it may fail, it should not crash the application at any
// given moment.
//
// Error: will throw LanguageTableMissing if the database is incorrectly set and the "languages" table is missing.
//
// Error: will throw LanguageParsingError if for some reason it is unable to parse the values it receives from the
// database.
//
// Error: will throw LanguageClosingTableError if it fails to close the database rows.
func (lr *LanguageRepository) GetLanguages() ([]models.Language, error) {
	query := "SELECT * FROM languages"
	rows, err := lr.connection.Query(query)

	if err != nil {
		apiErrors.CustomLog(err.Error(), apiErrors.ErrorLevel)
		return []models.Language{}, errors.New(apiErrors.LanguageTableMissing)
	}

	var languageList []models.Language
	var languageObj models.Language

	for rows.Next() {
		err = rows.Scan(
			&languageObj.Id,
			&languageObj.Name,
		)

		if err != nil {
			apiErrors.CustomLog(err.Error(), apiErrors.ErrorLevel)
			return []models.Language{}, errors.New(apiErrors.LanguageParsingError)
		}

		languageList = append(languageList, languageObj)
	}

	err = rows.Close()

	if err != nil {
		apiErrors.CustomLog(err.Error(), apiErrors.ErrorLevel)
		return []models.Language{}, errors.New(apiErrors.LanguageClosingTableError)
	}

	return languageList, nil
}

// GetLanguageById :
// Returns a "language" instance by id. Even though it may fail, it should not crash the application at any given
// moment.
//
// Error: will throw LanguageNotFound if a language with the provided id is not found.
func (lr *LanguageRepository) GetLanguageById(id int) (*models.Language, error) {
	query, err := lr.connection.Prepare("SELECT * FROM languages WHERE id = $1")

	if err != nil {
		apiErrors.CustomLog(err.Error(), apiErrors.ErrorLevel)
		return nil, errors.New(apiErrors.LanguageNotFound)
	}

	var languageObj models.Language
	err = query.QueryRow(id).Scan(&languageObj.Id, &languageObj.Name)

	if err != nil {
		apiErrors.CustomLog(err.Error(), apiErrors.ErrorLevel)
		return nil, errors.New(apiErrors.LanguageNotFound)
	}

	return &languageObj, nil
}

// GetLanguageByName :
// Returns a "language" instance by name. Even though it may fail, it should not crash the application at any given
// moment.
//
// Error: will throw LanguageNotFound if a language with the provided name is not found.
func (lr *LanguageRepository) GetLanguageByName(name string) (*models.Language, error) {
	query, err := lr.connection.Prepare("SELECT * FROM languages WHERE name = $1")

	if err != nil {
		customErrors.CustomLog(err.Error(), customErrors.ErrorLevel)
		return nil, err
	}

	var languageObj models.Language
	name = strings.ToLower(name)
	err = query.QueryRow(name).Scan(&languageObj.Id, &languageObj.Name)

	if err != nil {
		customErrors.CustomLog(err.Error(), customErrors.ErrorLevel)
		return nil, errors.New(customErrors.LanguageNotFound)
	}

	return &languageObj, nil
}
