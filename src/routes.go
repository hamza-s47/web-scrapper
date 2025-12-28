package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/help", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.POST("/urls", func(c *gin.Context) {
		var newUrls Urls

		if err := c.BindJSON(&newUrls); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, ScrapeService(newUrls))

	})

	r.Run()
}
