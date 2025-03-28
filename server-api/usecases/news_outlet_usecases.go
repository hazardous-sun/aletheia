package usecases

import (
	"ai-fact-checker/models"
	"ai-fact-checker/repositories"
)

type NewsOutletUseCase struct {
	newsOutletRepository repositories.NewsOutletRepository
}

func NewNewsOutletUsecase(repo repositories.NewsOutletRepository) NewsOutletUseCase {
	return NewsOutletUseCase{
		newsOutletRepository: repo,
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
func (no *NewsOutletUseCase) AddNewsOutlet(newsOutlet models.NewsOutlet) (models.NewsOutlet, error) {
	id, err := no.newsOutletRepository.AddNewsOutlet(newsOutlet)

	if err != nil && id < 0 {
		return models.NewsOutlet{}, err
	}

	return newsOutlet, nil
}
