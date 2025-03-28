package controllers

import (
	apiErrors "ai-fact-checker/errors"
	"ai-fact-checker/models"
	"ai-fact-checker/usecases"
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

// Create -----------------------------------------------------------------------------------------------------------------------------------

// Read -------------------------------------------------------------------------------------------------------------------------------------

// GetLanguages :
// Returns all the languages stored in the database. Even though it may fail, it should not crash the application at any given moment.
//
// Error: will return StatusBadRequest if the requet is invalid.
//
// Error: will return StatusInternalServerError if there's a problem while closing the "language" table or if it is missing.
func (lc *LanguageController) GetLanguages(ctx *gin.Context) {
	languages, err := lc.languageUseCase.GetLanguages()

	if err != nil {
		if err.Error() == apiErrors.LanguageParsingError {
			// LanguageParsingError
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
// Returns a "language" instance by id. Even though it may fail, it should not crash the application at any given moment.
//
// Error: will return StatusBadRequest if the request is invalid.
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
