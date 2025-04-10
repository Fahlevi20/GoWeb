package main

import (
	"context"
	"net/http"
	"time"

	"learn-sql-api/connection" // ga
	"learn-sql-api/model"

	"github.com/gin-gonic/gin"
)

func main() {
	connection.InitDB()

	router := gin.Default()

	router.POST("/query", func(c *gin.Context) {
		var req model.QueryRequest
		if err := c.BindJSON(&req); err != nil || req.Query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		data, err := connection.QueryData(ctx, req.Query) // <-- ini yang benar
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
