package main

import (
	"github.com/gin-gonic/gin"
	// "fmt"

	"github.com/hamza-s47/web-scrapper/src"
)

func main() {
	r := gin.Default()
	src.RegisterRoutes(r)
}
