package usecases

import (
	"ai-fact-checker/models"
	"ai-fact-checker/repositories"
)

type LanguageUseCase struct {
	languageRepository *repositories.LanguageRepository
}

func NewLanguageUsecase(repo *repositories.LanguageRepository) LanguageUseCase {
	return LanguageUseCase{
		languageRepository: repo,
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
func (lu LanguageUseCase) AddLanguage(language models.Language) (models.Language, error) {
	id, err := lu.languageRepository.AddLanguage(language)

	if err != nil && id > 0 {
		return models.Language{}, err
	}

	return language, nil
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
func (lu *LanguageUseCase) GetLanguages() ([]models.Language, error) {
	return lu.languageRepository.GetLanguages()
}

// GetLanguageById :
// Returns a "language" instance by id. Even though it may fail, it should not crash the application at any given
// moment.
//
// Error: will throw LanguageNotFound if a language with the provided id is not found.
func (lu *LanguageUseCase) GetLanguageById(id int) (*models.Language, error) {
	language, err := lu.languageRepository.GetLanguageById(id)

	if err != nil {
		return nil, err
	}

	return language, nil
}

// GetLanguageByName :
// Returns a "language" instance by name. Even though it may fail, it should not crash the application at any given
// moment.
//
// Error: will throw LanguageNotFound if a language with the provided name is not found.
func (lu *LanguageUseCase) GetLanguageByName(name string) (*models.Language, error) {
	language, err := lu.languageRepository.GetLanguageByName(name)

	if err != nil {
		return nil, err
	}

	return language, nil
}
