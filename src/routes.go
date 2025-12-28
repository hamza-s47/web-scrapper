package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/help", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"how_to_use": "Post the request body as an array of domains (use param '/urls') or by CLI",
			"body":       "{[\"http://url1.com\", \"http://url2.com\", \"http://url.com\"]}",
			"cli":        "go run main.go http://url1.com http://url2.com ... http://url.com",
			"help":       "go run main.go --help",
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
