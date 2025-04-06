package main

import (
	"aletheia-server/src/controllers"
	"aletheia-server/src/db"
	"aletheia-server/src/errors"
	"aletheia-server/src/repositories"
	"aletheia-server/src/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Database server
	dbConnection, err := db.ConnectDB()

	if err != nil {
		server_errors.Log(err.Error(), server_errors.ErrorLevel)
		return
	}

	// Initializing the repository layer
	languageRepository := repositories.NewLanguageRepository(dbConnection)
	languageUsecase := usecases.NewLanguageUsecase(languageRepository)
	languageController := controllers.NewLanguageController(languageUsecase)

	// Initializing the use case layer

	// Initializing the controller layer
	newsOutletRepository := repositories.NewNewsOutletRepository(dbConnection, *languageRepository)
	newsOutletUsecase := usecases.NewNewsOutletUsecase(newsOutletRepository)
	newsOutletController := controllers.NewNewsOutletController(newsOutletUsecase)

	// Initializing crawlers
	crawlerUsecase := usecases.NewCrawlerUsecase()
	crawlerController := controllers.NewCrawlerController(crawlerUsecase, newsOutletUsecase)

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

	err = server.Run(":8000")

	if err != nil {
		server_errors.Log(err.Error(), server_errors.ErrorLevel)
		return
	}
}
