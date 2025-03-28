package controllers

import (
	customErrors "ai-fact-checker/errors"
	"ai-fact-checker/models"
	"ai-fact-checker/usecases"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewsOutletController struct {
	newsOutletUsecase usecases.NewsOutletUseCase
}

func NewNewsOutletController(usecase usecases.NewsOutletUseCase) NewsOutletController {
	return NewsOutletController{
		newsOutletUsecase: usecase,
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
func (no *NewsOutletController) AddNewsOutlet(ctx *gin.Context) {
	var newsOutlet models.NewsOutlet
	err := ctx.BindJSON(&newsOutlet)

	if err != nil {
		customErrors.CustomLog(customErrors.NewsOutletParsingError, customErrors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	newsOutlet, err = no.newsOutletUsecase.AddNewsOutlet(newsOutlet)

	if err != nil {
		customErrors.CustomLog(customErrors.NewsOutletNotAdded, customErrors.ErrorLevel)
		if errors.Is(err, errors.New(customErrors.NewsOutletParsingError)) {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		} else if errors.Is(err, errors.New(customErrors.NewsOutletTableMissing)) {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		} else if errors.Is(err, errors.New(customErrors.NewsOutletClosingTableError)) {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}
		return
	}

	createdNewsOutlet, err := no.newsOutletUsecase.GetNewsOutletByName(newsOutlet.Name)

	if err != nil {
		customErrors.CustomLog(customErrors.NewsOutletNotAdded, customErrors.ErrorLevel)
		if errors.Is(err, errors.New(customErrors.NewsOutletNotFound)) {
			ctx.JSON(http.StatusNotFound, models.Response{
				Message: err.Error(),
				Status:  http.StatusNotFound,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, createdNewsOutlet)
}

// Read ----------------------------------------------------------------------------------------------------------------

// GetNewsOutletByName :
// Returns a "NewsOutlet" instance by name. Even though it may fail, it should not crash the application at any given
// moment.
//
// Error: will throw NewsOutletNotFound if a news outlet with the provided name is not found.
func (no *NewsOutletController) GetNewsOutletByName(ctx *gin.Context) {
	name := ctx.Param("newsOutletName")

	if name == "" {
		customErrors.CustomLog(customErrors.EmptyNameError, customErrors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: customErrors.EmptyNameError,
			Status:  http.StatusBadRequest,
		})
		return
	}

	newsOutlet, err := no.newsOutletUsecase.GetNewsOutletByName(name)

	if err != nil {
		customErrors.CustomLog(customErrors.NewsOutletNotFound, customErrors.ErrorLevel)
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		})
		return
	}

	ctx.JSON(http.StatusOK, newsOutlet)
}
