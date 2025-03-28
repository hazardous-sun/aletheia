package main

import (
	"ai-fact-checker/server-api/controller"
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
    db, err := db.ConnectDB()

    if err != nil {
        custom_errors.CustomLog(err.Error(), custom_errors.ErrorLevel)
        return
    }

    // Initializing the repository layer
    languageRepository := repositories.NewLanguageRepository(db)

    // Initializing the use case layer
    languageUsecase := usecases.NewLanguageUsecase(languageRepository)

    // Initializing the controller layer
    languageController := controllers.NewLanguageController(languageUsecase)

    // Initialize the API server
    server := gin.Default()

    // Setting up HTTP paths in the API server--------------------------------

    server.GET("/ping", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "message": "pong",
            "status": http.StatusOK,
        })
    })

    server.GET("languages", languageController.GetLanguages)
    server.GET("language/:languageId", languageController.GetLanguageById)

    // ------------------------------------------------------------------------
    
    err = server.Run(":8000")

    if err != nil {
        custom_errors.CustomLog(err.Error(), custom_errors.ErrorLevel)
        return
    }
}
