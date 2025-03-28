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

	if errors.Is(err, errors.New(customErrors.LanguageParsingError)) {
		ctx.JSON(http.StatusBadRequest, err)
	}

	language, err = lc.languageUseCase.AddLanguage(language)

	if err != nil {
		if errors.Is(err, errors.New(customErrors.LanguageParsingError)) {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			// LanguageTableMissing
			// LanguageClosingTableError
			ctx.JSON(http.StatusInternalServerError, err)
		}
		return
	}

	createdLanguage, err := lc.languageUseCase.GetLanguageByName(language.Name)

	if err != nil {
		customErrors.CustomLog(customErrors.LanguageNotAdded, customErrors.ErrorLevel)
		ctx.JSON(http.StatusInternalServerError, err)
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
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			// LanguageTableMissing
			// LanguageClosingTableError
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
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
		ctx.JSON(http.StatusBadRequest, models.Response{Message: apiErrors.EmptyIdError})
		return
	}

	languageId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: apiErrors.InvalidIdError})
		return
	}

	language, err := lc.languageUseCase.GetLanguageById(languageId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{Message: apiErrors.LanguageNotFound})
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
		ctx.JSON(http.StatusBadRequest, models.Response{Message: customErrors.EmptyIdError})
		return
	}

	language, err := lc.languageUseCase.GetLanguageByName(name)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{Message: customErrors.LanguageNotFound})
		return
	}

	ctx.JSON(http.StatusOK, language)
}
