package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "bienvenidos/as a go web :) \n",
		})
	})

	router.Run()
}
