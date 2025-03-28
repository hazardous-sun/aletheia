package usecases

import (
	custom_errors "ai-fact-checker/server-api/errors"
	"ai-fact-checker/server-api/models"
	"ai-fact-checker/server-api/repositories"
	"errors"
)

type LanguageUseCase struct {
    languageRepository repositories.LanguageRepository
}

func NewLanguageUsecase(repo repositories.LanguageRepository) LanguageUseCase {
	return LanguageUseCase{
		languageRepository: repo,
	}
}

// Read ------------------------------------------------------------------------------------------------------------------------------------

func (lu *LanguageUseCase) GetLanguages() ([]models.Language, error) {
    return lu.languageRepository.GetLanguages()
}

func (lu *LanguageUseCase) GetLanguageById(id int) (*models.Language, error) {
    language, err := lu.languageRepository.GetLanguageById(id)

	if err != nil {
		return nil, errors.New(custom_errors.LanguageNotFound)
	}

	return language, nil
}
