package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	recipeApi := g.Group("/api")
	{
		recipeApi.GET("/recipe", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "good")
		})
	}
	g.Run(":80")
}