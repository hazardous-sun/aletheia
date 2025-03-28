package main

import (
	controllers2 "fact-checker-server/src/controllers"
	"fact-checker-server/src/db"
	customErros "fact-checker-server/src/errors"
	repositories2 "fact-checker-server/src/repositories"
	usecases2 "fact-checker-server/src/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
) 

func main() {
	// Initialize the Database server
	dbConnection, err := db.ConnectDB()

    if err != nil {
        custom_errors.CustomLog(err.Error(), custom_errors.ErrorLevel)
        return
    }

	// Initializing the repository layer
	languageRepository := repositories2.NewLanguageRepository(dbConnection)
	languageUsecase := usecases2.NewLanguageUsecase(languageRepository)
	languageController := controllers2.NewLanguageController(languageUsecase)

	// Initializing the use case layer

	// Initializing the controller layer
	newsOutletRepository := repositories2.NewNewsOutletRepository(dbConnection, *languageRepository)
	newsOutletUsecase := usecases2.NewNewsOutletUsecase(newsOutletRepository)
	newsOutletController := controllers2.NewNewsOutletController(newsOutletUsecase)

	// Initializing crawlers
	crawlerUsecase := usecases2.NewCrawlerUsecase()
	crawlerController := controllers2.NewCrawlerController(crawlerUsecase, newsOutletUsecase)

	// Initialize the API server
	server := gin.Default()

	// Setting up HTTP paths in the API server -------------------------------------------------------------------------
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
			"status":  http.StatusOK,
		})
	})
	// ----- Languages
	// ---------- Create
	server.POST("language", languageController.AddLanguage)
	// ---------- Read
	server.GET("languages", languageController.GetLanguages)
	server.GET("languageId/:languageId", languageController.GetLanguageById)
	server.GET("languageName/:languageName", languageController.GetLanguageByName)
	// -----
	// ----- News Outlets
	// ---------- Create
	server.POST("newsOutlet", newsOutletController.AddNewsOutlet)
	// ---------- Read
	server.GET("newsOutlets", newsOutletController.GetNewsOutlets)
	server.GET("newsOutletName/:newsOutletName", newsOutletController.GetNewsOutletByName)
	server.GET("newsOutletId/:newsOutletId", newsOutletController.GetNewsOutletById)
	// ----- Crawlers
	server.POST("crawl", crawlerController.Crawl)
	// -----------------------------------------------------------------------------------------------------------------

    if err != nil {
        custom_errors.CustomLog(err.Error(), custom_errors.ErrorLevel)
        return
    }
}
