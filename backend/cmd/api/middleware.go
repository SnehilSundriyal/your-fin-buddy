package main

import "github.com/gin-gonic/gin"

func (app *application) enableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set() == c.Header()
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
