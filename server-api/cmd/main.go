package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
) 

func main() {
    // Initialize the server
    server := gin.Default()

    // Setting up HTTP calls handling

    server.GET("/ping", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "message": "pong",
            "status": http.StatusOK,
        })
    })

    // ------------------------------------------------------------------------
    
    err := server.Run(":8000")

    if err != nil {
        panic(err)
    }
}
