package controllers

import (
	apiErrors "ai-fact-checker/errors"
	"ai-fact-checker/models"
	"ai-fact-checker/usecases"
	"errors"
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
	var language models.Language
	err := ctx.BindJSON(&language)

	if err != nil {
		customErrors.CustomLog(customErrors.LanguageParsingError, customErrors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	language, err = lc.languageUseCase.AddLanguage(language)

	if err != nil {
		customErrors.CustomLog(customErrors.LanguageNotAdded, customErrors.ErrorLevel)
		if errors.Is(err, errors.New(customErrors.LanguageParsingError)) {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		} else if errors.Is(err, errors.New(customErrors.LanguageTableMissing)) {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		} else if errors.Is(err, errors.New(customErrors.LanguageClosingTableError)) {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}
		return
	}

	createdLanguage, err := lc.languageUseCase.GetLanguageByName(language.Name)

	if err != nil {
		customErrors.CustomLog(customErrors.LanguageNotAdded, customErrors.ErrorLevel)
		if errors.Is(err, errors.New(customErrors.LanguageNotFound)) {
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

	ctx.JSON(http.StatusOK, createdLanguage)
}

// Read ----------------------------------------------------------------------------------------------------------------

// GetLanguages :
// Returns all the languages stored in the database. Even though it may fail, it should not crash the application at any
// given moment.
//
// Error: will return StatusBadRequest if the requet is invalid.
//
// Error: will return StatusInternalServerError if there's a problem while closing the "language" table or if it is
// missing.
func (lc *LanguageController) GetLanguages(ctx *gin.Context) {
	languages, err := lc.languageUseCase.GetLanguages()

	if err != nil {
		if errors.Is(err, errors.New(customErrors.LanguageParsingError)) {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		} else if errors.Is(err, errors.New(customErrors.LanguageTableMissing)) {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		} else if errors.Is(err, errors.New(customErrors.LanguageClosingTableError)) {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, models.Response{
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
		customErrors.CustomLog(customErrors.EmptyIdError, customErrors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: customErrors.EmptyIdError,
			Status:  http.StatusBadRequest,
		})
		return
	}

	languageId, err := strconv.Atoi(id)

	if err != nil {
		customErrors.CustomLog(customErrors.InvalidIdError, customErrors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: customErrors.InvalidIdError,
			Status:  http.StatusBadRequest,
		})
		return
	}

	language, err := lc.languageUseCase.GetLanguageById(languageId)

	if err != nil {
		customErrors.CustomLog(customErrors.LanguageNotFound, customErrors.ErrorLevel)
		ctx.JSON(http.StatusNotFound, models.Response{
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
		customErrors.CustomLog(customErrors.EmptyNameError, customErrors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: customErrors.EmptyNameError,
			Status:  http.StatusBadRequest,
		})
		return
	}

	language, err := lc.languageUseCase.GetLanguageByName(name)

	if err != nil {
		customErrors.CustomLog(customErrors.LanguageNotFound, customErrors.ErrorLevel)
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		})
		return
	}

	ctx.JSON(http.StatusOK, language)
}
