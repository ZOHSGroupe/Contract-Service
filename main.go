package main

import (
	"Contract-Service/app/configs"
	"Contract-Service/app/middlewares"
	"Contract-Service/app/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()
	//run database
	configs.ConnectDB()
	// Use the token verification middleware for routes that require authentication
	router.Use(middlewares.VerifyToken())
	//routes
	routes.ContractRoute(router) //add this
	router.Run("localhost:" + os.Getenv("GO_DOCKER_PORT"))

}
