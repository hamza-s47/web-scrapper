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
			return
		}

		c.IndentedJSON(http.StatusCreated, newUrls)
	})

	r.Run()
}
