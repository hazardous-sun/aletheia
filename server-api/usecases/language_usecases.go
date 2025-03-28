package usecases

import (
	"ai-fact-checker/server-api/models"
	"ai-fact-checker/server-api/repositories"
)

type LanguageUseCase struct {
    languageRepository repositories.LanguageRepository
}

func NewLanguageUsecase(repo repositories.LanguageRepository) LanguageUseCase {
	return LanguageUseCase{
		languageRepository: repo,
	}
}

func (lu *LanguageUseCase) GetLanguages() ([]models.Language, error) {
    return lu.languageRepository.GetLanguages()
}