package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

const PORT = ":8080"

type application struct {
}

func main() {
	var app application

	gin.SetMode(gin.ReleaseMode)

	server := app.routes()

	err := server.Run(PORT)
	if err != nil {
		log.Fatal(err)
	}

}
