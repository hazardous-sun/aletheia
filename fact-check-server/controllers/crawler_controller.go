package controllers

import (
	customErrors "ai-fact-checker/errors"
	"ai-fact-checker/models"
	"ai-fact-checker/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CrawlerController struct {
	crawlerUseCase    usecases.CrawlerUsecase
	newsOutletUseCase usecases.NewsOutletUseCase
}

func NewCrawlerController(crawler usecases.CrawlerUsecase) CrawlerController {
	return CrawlerController{
		crawlerUseCase: crawler,
	}
}

func (cr *CrawlerController) Crawl(ctx *gin.Context) {
	var crawlersInitializer models.CrawlerInitializer
	err := ctx.BindJSON(&crawlersInitializer)

	if err != nil {
		customErrors.CustomLog(customErrors.InvalidParameters, customErrors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	newsOutlets, err := cr.newsOutletUseCase.GetNewsOutlets()

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

	cr.crawlerUseCase.Crawl(newsOutlets, crawlersInitializer.PagesToVisit, crawlersInitializer.Query)
}
