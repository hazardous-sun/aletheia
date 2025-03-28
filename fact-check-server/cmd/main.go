package main

import (
	"ai-fact-checker/server-api/controllers"
	"ai-fact-checker/server-api/db"
	custom_errors "ai-fact-checker/server-api/errors"
	"ai-fact-checker/server-api/repositories"
	"ai-fact-checker/server-api/usecases"
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
	languageRepository := repositories.NewLanguageRepository(dbConnection)
	languageUsecase := usecases.NewLanguageUsecase(languageRepository)
	languageController := controllers.NewLanguageController(languageUsecase)

	// Initializing the use case layer

	// Initializing the controller layer
	newsOutletRepository := repositories.NewNewsOutletRepository(dbConnection, *languageRepository)
	newsOutletUsecase := usecases.NewNewsOutletUsecase(newsOutletRepository)
	newsOutletController := controllers.NewNewsOutletController(newsOutletUsecase)

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
	// -----------------------------------------------------------------------------------------------------------------

    if err != nil {
        custom_errors.CustomLog(err.Error(), custom_errors.ErrorLevel)
        return
    }
}
