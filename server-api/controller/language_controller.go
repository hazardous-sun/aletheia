package controller

import (
	"github.com/gin-gonic/gin"
	"ai-fact-checker/server-api/usecases"
	"net/http"
)

type LanguageController struct {
	languageUseCase usecases.LanguageUseCase
}

func NewLanguageController(usecase usecases.LanguageUseCase) LanguageController {
	return LanguageController{
		languageUseCase: usecase,
	}
}

func (lc *LanguageController) GetLanguages(ctx *gin.Context) {
	languages, err := lc.languageUseCase.GetLanguages()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, languages)
}