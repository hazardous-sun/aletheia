package controllers

import (
	"fact-checker-server/src/errors"
	models2 "fact-checker-server/src/models"
	usecases2 "fact-checker-server/src/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CrawlerController struct {
	crawlerUseCase    usecases2.CrawlerUsecase
	newsOutletUseCase usecases2.NewsOutletUseCase
}

func NewCrawlerController(crawler usecases2.CrawlerUsecase, newsOutletUseCase usecases2.NewsOutletUseCase) CrawlerController {
	return CrawlerController{
		crawlerUseCase:    crawler,
		newsOutletUseCase: newsOutletUseCase,
	}
}

func (cr *CrawlerController) Crawl(ctx *gin.Context) {
	var crawlersInitializer models2.CrawlerInitializer
	err := ctx.BindJSON(&crawlersInitializer)

	if err != nil {
		custom_errors.Log(custom_errors.InvalidParameters, custom_errors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models2.Response{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	newsOutlets, err := cr.newsOutletUseCase.GetNewsOutlets()

	if err != nil {
		custom_errors.Log(err.Error(), custom_errors.ErrorLevel)
		switch err.Error() {
		case custom_errors.NewsOutletParsingError:
			ctx.JSON(http.StatusBadRequest, models2.Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		case custom_errors.NewsOutletTableMissing:
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

	cr.crawlerUseCase.Crawl(newsOutlets, crawlersInitializer.PagesToVisit, crawlersInitializer.Query)
}
