package main

import (
	"cleanArchGo/infrastructure/databases"
	"cleanArchGo/infrastructure/environments"
	"cleanArchGo/infrastructure/routers"
	"os"
)

func init() {
	environments.LoadEnv()
	databases.NewMySQLConnection()
}

func main() {
	router := routers.SetupRouter()
	router.Run(os.Getenv("APP_PORT"))
}
