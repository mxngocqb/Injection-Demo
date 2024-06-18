package main

import (
	"github.com/gin-gonic/gin"
	_ "injection_testing/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(c *Controllers) *gin.Engine {
	router := gin.Default()

	drivers := router.Group("/drivers")
	{
		drivers.POST("", c.DriverController.CreateDriver)
		drivers.GET("", c.DriverController.ReadAllDriver)
		drivers.GET("/:driverID", c.DriverController.ReadDriverByID)
		drivers.PUT("/:driverID", c.DriverController.UpdateDriver)
		drivers.DELETE("/:driverID", c.DriverController.DeleteDriver)
	}

	router.GET("/list",  c.DriverController.List)
	router.GET("/show",  c.DriverController.CodeInjection)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
