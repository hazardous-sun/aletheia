package controllers

import (
	"fact-checker-server/src/errors"
	models2 "fact-checker-server/src/models"
	"fact-checker-server/src/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LanguageController struct {
	languageUseCase usecases.LanguageUseCase
}

func NewLanguageController(usecase usecases.LanguageUseCase) LanguageController {
	return LanguageController{
		languageUseCase: usecase,
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
func (lc *LanguageController) AddLanguage(ctx *gin.Context) {
	var language models2.Language
	err := ctx.BindJSON(&language)

	if err != nil {
		custom_errors.CustomLog(custom_errors.LanguageParsingError, custom_errors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models2.Response{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	language, err = lc.languageUseCase.AddLanguage(language)

	if err != nil {
		custom_errors.CustomLog(custom_errors.LanguageNotAdded, custom_errors.ErrorLevel)
		switch err.Error() {
		case custom_errors.LanguageParsingError:
			ctx.JSON(http.StatusBadRequest, models2.Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		case custom_errors.LanguageTableMissing:
			ctx.JSON(http.StatusInternalServerError, models2.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		case custom_errors.LanguageClosingTableError:
			ctx.JSON(http.StatusInternalServerError, models2.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		case custom_errors.LanguageAlreadyExists:
			ctx.JSON(http.StatusBadRequest, models2.Response{
				Message: custom_errors.LanguageAlreadyExists,
				Status:  http.StatusBadRequest,
			})
		default:
			ctx.JSON(http.StatusInternalServerError, models2.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}
		return
	}

	createdLanguage, err := lc.languageUseCase.GetLanguageByName(language.Name)

	if err != nil {
		custom_errors.CustomLog(custom_errors.LanguageNotAdded, custom_errors.ErrorLevel)
		switch err.Error() {
		case custom_errors.LanguageNotFound:
			ctx.JSON(http.StatusNotFound, models2.Response{
				Message: err.Error(),
				Status:  http.StatusNotFound,
			})
		default:
			ctx.JSON(http.StatusInternalServerError, models2.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, createdLanguage)
}

// Read ----------------------------------------------------------------------------------------------------------------

// GetLanguages :
// Returns all the languages stored in the database. Even though it may fail, it should not crash the application at any
// given moment.
//
// Error: will return StatusBadRequest if the body is invalid.
//
// Error: will return StatusInternalServerError if there's a problem while closing the "language" table or if it is
// missing.
func (lc *LanguageController) GetLanguages(ctx *gin.Context) {
	languages, err := lc.languageUseCase.GetLanguages()

	if err != nil {
		switch err.Error() {
		case custom_errors.LanguageParsingError:
			ctx.JSON(http.StatusBadRequest, models2.Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		case custom_errors.LanguageTableMissing:
			ctx.JSON(http.StatusInternalServerError, models2.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		case custom_errors.LanguageClosingTableError:
			ctx.JSON(http.StatusInternalServerError, models2.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		default:
			ctx.JSON(http.StatusInternalServerError, models2.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, languages)
}

// GetLanguageById :
// Returns a "language" instance by id. Even though it may fail, it should not crash the application at any given
// moment.
//
// Error: will return StatusBadRequest if the body is invalid.
//
// Error: will return StatusNotFound if a language with the provided id is not found.
func (lc *LanguageController) GetLanguageById(ctx *gin.Context) {
	id := ctx.Param("languageId")

	if id == "" {
		custom_errors.CustomLog(custom_errors.EmptyIdError, custom_errors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models2.Response{
			Message: custom_errors.EmptyIdError,
			Status:  http.StatusBadRequest,
		})
		return
	}

	languageId, err := strconv.Atoi(id)

	if err != nil {
		custom_errors.CustomLog(custom_errors.InvalidIdError, custom_errors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models2.Response{
			Message: custom_errors.InvalidIdError,
			Status:  http.StatusBadRequest,
		})
		return
	}

	language, err := lc.languageUseCase.GetLanguageById(languageId)

	if err != nil {
		custom_errors.CustomLog(custom_errors.LanguageNotFound, custom_errors.ErrorLevel)
		ctx.JSON(http.StatusNotFound, models2.Response{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		})
		return
	}

	ctx.JSON(http.StatusOK, language)
}

// GetLanguageByName :
// Returns a "language" instance by name. Even though it may fail, it should not crash the application at any given
// moment.
//
// Error: will return StatusBadRequest if the body is invalid.
//
// Error: will return StatusNotFound if a language with the provided name is not found.
func (lc *LanguageController) GetLanguageByName(ctx *gin.Context) {
	name := ctx.Param("languageName")

	if name == "" {
		custom_errors.CustomLog(custom_errors.EmptyNameError, custom_errors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models2.Response{
			Message: custom_errors.EmptyNameError,
			Status:  http.StatusBadRequest,
		})
		return
	}

	language, err := lc.languageUseCase.GetLanguageByName(name)

	if err != nil {
		custom_errors.CustomLog(custom_errors.LanguageNotFound, custom_errors.ErrorLevel)
		ctx.JSON(http.StatusNotFound, models2.Response{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		})
		return
	}

	ctx.JSON(http.StatusOK, language)
}
