package controllers

import (
	customErrors "ai-fact-checker/errors"
	"ai-fact-checker/models"
	"ai-fact-checker/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
// Error: will throw LanguageTableMissing if the database is incorrectly set and the "languages" table is missing.
//
// Error: will throw LanguageParsingError if for some reason it is unable to parse the values it receives from the
// database.
//
// Error: will throw LanguageClosingTableError if it fails to close the database rows.
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
		switch err.Error() {
		case customErrors.LanguageParsingError:
			ctx.JSON(http.StatusBadRequest, models.Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		case customErrors.NewsOutletTableMissing:
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		case customErrors.NewsOutletClosingTableError:
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		case customErrors.LanguageNotFound:
			ctx.JSON(http.StatusNotFound, models.Response{
				Message: err.Error(),
				Status:  http.StatusNotFound,
			})
		default:
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}
		return
	}

	// Confirm it was correctly added
	createdNewsOutlet, err := no.newsOutletUsecase.GetNewsOutletByName(newsOutlet.Name)

	if err != nil {
		customErrors.CustomLog(customErrors.NewsOutletNotAdded, customErrors.ErrorLevel)
		switch err.Error() {
		case customErrors.NewsOutletNotFound:
			ctx.JSON(http.StatusNotFound, models.Response{
				Message: err.Error(),
				Status:  http.StatusNotFound,
			})
		default:
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

// GetNewsOutlets :
// Returns all the news outlets stored in the database. Even though it may fail, it should not crash the application at
// any given moment.
//
// Error: will return StatusBadRequest if the body is invalid.
//
// Error: will return StatusInternalServerError if there's a problem while closing the "news_outlet" table or if it is
// missing.
func (no *NewsOutletController) GetNewsOutlets(ctx *gin.Context) {
	newsOutlets, err := no.newsOutletUsecase.GetNewsOutlets()

	if err != nil {
		customErrors.CustomLog(err.Error(), customErrors.ErrorLevel)
		switch err.Error() {
		case customErrors.NewsOutletParsingError:
			ctx.JSON(http.StatusBadRequest, models.Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		case customErrors.NewsOutletTableMissing:
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		case customErrors.LanguageClosingTableError:
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		default:
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, newsOutlets)
}

// GetNewsOutletByName :
// Returns a "NewsOutlet" instance by name. Even though it may fail, it should not crash the application at any given
// moment.
//
// Error: will return StatusBadRequest if the body is invalid.
//
// Error: will return StatusNotFound if a news outlet with the provided name is not found.
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

// GetNewsOutletById :
// Returns a "NewsOutlet" instance by id. Even though it may fail, it should not crash the application at any given
// moment.
//
// Error: will return StatusBadRequest if the body is invalid.
//
// Error: will return StatusNotFound if a news outlet with the provided name is not found.
func (no *NewsOutletController) GetNewsOutletById(ctx *gin.Context) {
	inputId := ctx.Param("newsOutletName")

	if inputId == "" {
		customErrors.CustomLog(customErrors.EmptyNameError, customErrors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: customErrors.EmptyNameError,
			Status:  http.StatusBadRequest,
		})
		return
	}

	id, err := strconv.Atoi(inputId)

	if err != nil {
		customErrors.CustomLog(customErrors.InvalidIdError, customErrors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	newsOutlet, err := no.newsOutletUsecase.GetNewsOutletById(id)

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
