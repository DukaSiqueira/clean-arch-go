package main

import (
	"cleanArchGo/infrastructure/environments"
	"cleanArchGo/infrastructure/databases"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	environments.LoadEnv()
	databases.NewMySQLConnection()

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!!",
		})
	})

	router.Run(os.Getenv("APP_PORT"))
}
