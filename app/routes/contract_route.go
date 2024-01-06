package routes

import (
	"Contract-Service/app/controllers"

	"github.com/gin-gonic/gin"
)

func ContractRoute(router *gin.Engine) {
	router.POST("/contract", controllers.CreateContract())
	router.GET("/contract", controllers.GetAllContracts())
	router.GET("/contract/:id", controllers.GetAContract())
	router.DELETE("/contract/:id", controllers.DeleteAContract())
	router.GET("/contract/client/:nationalId", controllers.GetContractsByNationalID())
	router.GET("/contract/vihecule/:identificationNumber", controllers.GetContractsByVehicleIdentificationNumber())
	router.GET("/contract/validity/:identificationNumber", controllers.CheckContractValidityByViheculeIdentificationNumber())
	router.GET("/contract/valid/:identificationNumber", controllers.GetValiditContractOfViheculeByIdentificationNumber())
}
