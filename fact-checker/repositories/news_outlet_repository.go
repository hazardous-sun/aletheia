package repositories

import (
	"database/sql"
	"errors"
	customErrors "fact-checker-server/errors"
	"fact-checker-server/models"
	"strings"
)

type NewsOutletRepository struct {
	connection         *sql.DB
	languageRepository *LanguageRepository
}

func NewNewsOutletRepository(connection *sql.DB, languageRepository LanguageRepository) NewsOutletRepository {
	return NewsOutletRepository{
		connection:         connection,
		languageRepository: &languageRepository,
	}
}

// Create --------------------------------------------------------------------------------------------------------------

// AddNewsOutlet :
// Creates a new news outlet inside the database based on the model received as parameter.
//
// Error: will throw LanguageNotFound if the provided language is not maintained inside the database.
//
// Error: will throw NewsOutletTableMissing if the database is incorrectly set and the "news_outlet" table is missing.
//
// Error: will throw NewsOutletParsingError if for some reason it is unable to parse the values it receives from the
// database.
//
// Error: will throw NewsOutletClosingTableError if it fails to close the database rows.
func (no *NewsOutletRepository) AddNewsOutlet(newsOutlet models.NewsOutlet) (int, error) {
	// Collect newsOutlet.Language id inside the languages table
	language, err := no.languageRepository.GetLanguageByName(newsOutlet.Language)

	if err != nil {
		customErrors.CustomLog(customErrors.LanguageNotFound, customErrors.ErrorLevel)
		if errors.Is(err, sql.ErrNoRows) {
			return -1, errors.New(customErrors.NewsOutletNotFound)
		}
		return -1, err
	}

	languageId := language.Id

	// Insert newsOutlet into the database
	query, err := no.connection.Prepare("INSERT INTO news_outlet (name, queryurl, htmlselector, languageid, credibility) VALUES ($1, $2, $3, $4, $5) RETURNING id")

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			customErrors.CustomLog(customErrors.NewsOutletTableMissing, customErrors.ErrorLevel)
			return -1, errors.New(customErrors.NewsOutletTableMissing)
		}
		return -1, err
	}

	var id int
	name := strings.ToLower(newsOutlet.Name)
	err = query.QueryRow(name, newsOutlet.QueryUrl, newsOutlet.HtmlSelector, languageId, newsOutlet.Credibility).Scan(&id)

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

// GetNewsOutlets :
// Returns all the news outlets stored in the database. Even though it may fail, it should not crash the application at
// any given moment.
//
// Error: will throw NewsOutletTableMissing if the database is incorrectly set and the "news_outlet" table is missing.
//
// Error: will throw NewsOutletParsingError if for some reason it is unable to parse the values it receives from the
// database.
//
// Error: will throw NewsOutletClosingTableError if it fails to close the database rows.
func (no *NewsOutletRepository) GetNewsOutlets() ([]models.NewsOutlet, error) {
	query := "SELECT * FROM news_outlet"
	rows, err := no.connection.Query(query)

	if err != nil {
		customErrors.CustomLog(err.Error(), customErrors.ErrorLevel)
		return []models.NewsOutlet{}, errors.New(customErrors.NewsOutletTableMissing)
	}

	var newsOutletList []models.NewsOutlet
	var newsOutletObj models.NewsOutlet
	var languageId int

	for rows.Next() {
		err = rows.Scan(
			&newsOutletObj.Id,
			&newsOutletObj.Name,
			&newsOutletObj.QueryUrl,
			&newsOutletObj.HtmlSelector,
			&languageId,
			&newsOutletObj.Credibility,
		)

		if err != nil {
			customErrors.CustomLog(err.Error(), customErrors.ErrorLevel)
			return []models.NewsOutlet{}, errors.New(customErrors.NewsOutletParsingError)
		}

		language, err := no.languageRepository.GetLanguageById(languageId)

		if err != nil {
			customErrors.CustomLog(customErrors.LanguageNotFound, customErrors.ErrorLevel)
			return []models.NewsOutlet{}, err
		}

		newsOutletObj.Language = language.Name
		newsOutletList = append(newsOutletList, newsOutletObj)
	}

	err = rows.Close()

	if err != nil {
		customErrors.CustomLog(err.Error(), customErrors.ErrorLevel)
		return []models.NewsOutlet{}, errors.New(customErrors.NewsOutletClosingTableError)
	}

	return newsOutletList, nil
}

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
	var languageId int
	name = strings.ToLower(name)
	err = query.QueryRow(name).Scan(&newsOutletObj.Id, &newsOutletObj.Name, &newsOutletObj.QueryUrl, &newsOutletObj.HtmlSelector, &languageId, &newsOutletObj.Credibility)

	if err != nil {
		customErrors.CustomLog(err.Error(), customErrors.ErrorLevel)
		return nil, errors.New(customErrors.NewsOutletNotFound)
	}

	language, err := no.languageRepository.GetLanguageById(languageId)

	if err != nil {
		customErrors.CustomLog(customErrors.LanguageNotFound, customErrors.ErrorLevel)
		return &models.NewsOutlet{}, err
	}

	newsOutletObj.Language = language.Name

	return &newsOutletObj, nil
}

// GetNewsOutletById :
// Returns a "NewsOutlet" instance by id. Even though it may fail, it should not crash the application at any given
// moment.
//
// Error: will throw NewsOutletNotFound if a news outlet with the provided id is not found.
func (no *NewsOutletRepository) GetNewsOutletById(id int) (*models.NewsOutlet, error) {
	query, err := no.connection.Prepare("SELECT * FROM news_outlet WHERE id = $1")

	if err != nil {
		customErrors.CustomLog(err.Error(), customErrors.ErrorLevel)
		return nil, err
	}

	var newsOutletObj models.NewsOutlet
	var languageId int
	err = query.QueryRow(id).Scan(&newsOutletObj.Id, &newsOutletObj.Name, &newsOutletObj.QueryUrl, &newsOutletObj.HtmlSelector, &languageId, &newsOutletObj.Credibility)

	if err != nil {
		customErrors.CustomLog(err.Error(), customErrors.ErrorLevel)
		return nil, errors.New(customErrors.NewsOutletNotFound)
	}

	language, err := no.languageRepository.GetLanguageById(languageId)

	if err != nil {
		customErrors.CustomLog(customErrors.LanguageNotFound, customErrors.ErrorLevel)
		return &models.NewsOutlet{}, err
	}

	newsOutletObj.Language = language.Name

	return &newsOutletObj, nil
}
