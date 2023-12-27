package main

import (
	"Contract-Service/app/configs"
	"Contract-Service/app/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()
	//run database
	configs.ConnectDB()
	//routes
	routes.UserRoute(router) //add this
	router.Run("localhost:" + os.Getenv("GO_DOCKER_PORT"))

}
