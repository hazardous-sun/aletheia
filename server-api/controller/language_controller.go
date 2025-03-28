package controller

import (
	custom_errors "ai-fact-checker/server-api/errors"
	"ai-fact-checker/server-api/models"
	"ai-fact-checker/server-api/usecases"
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

func (lc *LanguageController) GetLanguages(ctx *gin.Context) {
	languages, err := lc.languageUseCase.GetLanguages()

	if err != nil {
		if err.Error() == custom_errors.LanguageParsingError {
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

func (lc *LanguageController) GetLanguageById(ctx *gin.Context) {
	id := ctx.Param("languageId")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: custom_errors.EmptyIdError})
		return
	}

	languageId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: custom_errors.InvalidIdError})
		return
	}

	language, err := lc.languageUseCase.GetLanguageById(languageId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{Message: custom_errors.LanguageNotFound})
		return
	}

	ctx.JSON(http.StatusOK, language)
}