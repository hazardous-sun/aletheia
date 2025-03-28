package controller

import (
	custom_errors "ai-fact-checker/server-api/errors"
	"ai-fact-checker/server-api/usecases"
	"net/http"

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