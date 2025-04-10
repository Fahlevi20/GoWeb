package main

import (
	"context"
	"net/http"
	"time"

	"learn-sql-api/connection" // ga

	"github.com/gin-gonic/gin"
)

func main() {
	connection.InitDB()

	router := gin.Default()

	router.GET("/query", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		data, err := connection.QueryData(ctx) // <-- ini yang benar
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	router.Run()
}
