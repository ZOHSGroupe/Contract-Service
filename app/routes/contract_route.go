package routes

import (
	"Contract-Service/app/controllers"

	"github.com/gin-gonic/gin"
)

func ContractRoute(router *gin.Engine) {
	router.POST("/contract", controllers.CreateContract())
	router.GET("/contract", controllers.GetAllContracts())
	router.GET("/contract/:id", controllers.GetAContract())
	router.GET("/contract/client/:nationalId", controllers.GetContractsByNationalID())
	router.GET("/contract/vihecule/:identificationNumber", controllers.GetContractsByVehicleIdentificationNumber())
	router.GET("/contract/validity/:identificationNumber", controllers.CheckContractValidityByViheculeIdentificationNumber())
}
