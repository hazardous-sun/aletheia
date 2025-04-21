package controllers

import (
	"aletheia-server/src/errors"
	"aletheia-server/src/models"
	"aletheia-server/src/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CrawlerController struct {
	crawlerUseCase    usecases.CrawlerUsecase
	newsOutletUseCase usecases.NewsOutletUseCase
}

func NewCrawlerController(crawler usecases.CrawlerUsecase, newsOutletUseCase usecases.NewsOutletUseCase) CrawlerController {
	return CrawlerController{
		crawlerUseCase:    crawler,
		newsOutletUseCase: newsOutletUseCase,
	}
}

func (cr *CrawlerController) Crawl(ctx *gin.Context) {
	var crawlersInitializer models.CrawlerInitializer
	err := ctx.BindJSON(&crawlersInitializer)

	if err != nil {
		server_errors.Log(server_errors.InvalidParameters, server_errors.ErrorLevel)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	newsOutlets, err := cr.newsOutletUseCase.GetNewsOutlets()

	if err != nil {
		server_errors.Log(err.Error(), server_errors.ErrorLevel)
		switch err.Error() {
		case server_errors.NewsOutletParsingError:
			ctx.JSON(http.StatusBadRequest, models.Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
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
