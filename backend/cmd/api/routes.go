package main

import "github.com/gin-gonic/gin"

func (app *application) routes() *gin.Engine {
	router := gin.Default()
	router.Use(app.enableCORS())

	router.GET("/", Home)

	return router
}
