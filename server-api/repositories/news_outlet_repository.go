package repositories

import (
	customErrors "ai-fact-checker/errors"
	"ai-fact-checker/models"
	"database/sql"
	"errors"
)

type NewsOutletRepository struct {
	connection *sql.DB
}

func NewNewsOutletRepository(connection *sql.DB) NewsOutletRepository {
	return NewsOutletRepository{
		connection: connection,
	}
}

// Create --------------------------------------------------------------------------------------------------------------

// AddNewsOutlet :
// Creates a new news outlet inside the database based on the model received as parameter.
//
// Error: will throw NewsOutletTableMissing if the database is incorrectly set and the "news_outlet" table is missing.
//
// Error: will throw NewsOutletParsingError if for some reason it is unable to parse the values it receives from the
// database.
//
// Error: will throw NewsOutletClosingTableError if it fails to close the database rows.
func (no *NewsOutletRepository) AddNewsOutlet(newsOutlet models.NewsOutlet) (int, error) {
	query, err := no.connection.Prepare("INSERT INTO news_outlet (name, url, language) VALUES ($1, $2, $3) RETURNING id")

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			customErrors.CustomLog(customErrors.NewsOutletTableMissing, customErrors.ErrorLevel)
			return -1, errors.New(customErrors.NewsOutletTableMissing)
		}
		return -1, err
	}

	var id int
	err = query.QueryRow(newsOutlet.Name, newsOutlet.Url, newsOutlet.Language).Scan(&id)

	if err != nil {
		customErrors.CustomLog(customErrors.NewsOutletParsingError, customErrors.ErrorLevel)
		return -1, errors.New(customErrors.NewsOutletParsingError)
	}

	err = query.Close()

	if err != nil {
		customErrors.CustomLog(customErrors.NewsOutletClosingTableError, customErrors.ErrorLevel)
		return -1, errors.New(customErrors.NewsOutletClosingTableError)
	}

	return id, nil
}

// Read ----------------------------------------------------------------------------------------------------------------

// GetNewsOutletByName :
// Returns a "NewsOutlet" instance by name. Even though it may fail, it should not crash the application at any given
// moment.
//
// Error: will throw NewsOutletNotFound if a news outlet with the provided name is not found.
func (no *NewsOutletRepository) GetNewsOutletByName(name string) (*models.NewsOutlet, error) {
	query, err := no.connection.Prepare("SELECT * FROM news_outlet WHERE name = $1")

	if err != nil {
		customErrors.CustomLog(err.Error(), customErrors.ErrorLevel)
		return nil, err
	}

	var newsOutletObj models.NewsOutlet
	err = query.QueryRow(name).Scan(&newsOutletObj.Id, &newsOutletObj.Name, &newsOutletObj.Url, &newsOutletObj.Language)

	if err != nil {
		customErrors.CustomLog(err.Error(), customErrors.ErrorLevel)
		return nil, errors.New(customErrors.NewsOutletNotFound)
	}

	return &newsOutletObj, nil
}
