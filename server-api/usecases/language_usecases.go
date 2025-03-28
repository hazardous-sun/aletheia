package usecases

import (
	"ai-fact-checker/server-api/models"
	"ai-fact-checker/server-api/repositories"
)

type LanguageUseCase struct {
    languageRepository *repositories.LanguageRepository
}

func NewLanguageUsecase(repo *repositories.LanguageRepository) LanguageUseCase {
	return LanguageUseCase{
		languageRepository: repo,
	}
}

// Read ------------------------------------------------------------------------------------------------------------------------------------

// GetLanguages :
// Returns all the languages stored in the database. Even though it may fail, it should not crash the application at any given moment.
//
// Error: will throw LanguageTableMissing if the database is incorrectly set and the "languages" table is missing.
//
// Error: will throw LanguageParsingError if for some reason it is unable to parse the values it receives from the database.
//
// Error: will throw LanguageClosingTableError if it fails to close the database rows.
func (lu *LanguageUseCase) GetLanguages() ([]models.Language, error) {
    return lu.languageRepository.GetLanguages()
}

// GetLanguageById :
// Returns a "language" instance by id. Even though it may fail, it should not crash the application at any given moment.
//
// Error: will throw LanguageNotFound if a language with the provided id is not found.
func (lu *LanguageUseCase) GetLanguageById(id int) (*models.Language, error) {
    language, err := lu.languageRepository.GetLanguageById(id)

	if err != nil {
		return nil, err
	}

	return language, nil
}
