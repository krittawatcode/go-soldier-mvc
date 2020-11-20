package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-soldier-mvc/controller"
	"github.com/krittawatcode/go-soldier-mvc/service"
)

func main() {
	server := gin.Default()
	v1 := server.Group("/v1")
	{
		v1.POST("/eat/:commission", func(c *gin.Context) {
			var dutyService service.DutyService = service.SoldierDutyService(c)
			var soldierController controller.SoldierController = controller.SoldierHandler(&dutyService)
			soldierController.Eat(c)
		})
	}
	port := "8080"
	server.Run(":" + port)
}
