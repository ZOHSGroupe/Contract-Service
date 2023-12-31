package routes

import (
	"Contract-Service/app/controllers"
	"github.com/gin-gonic/gin"
)

func ContractRoute(router *gin.Engine) {
	router.POST("/contract", controllers.CreateContract())
	router.GET("/contract/:id", controllers.GetAContract())
	router.DELETE("/contract/:id", controllers.DeleteAContract())
}
