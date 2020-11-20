package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-testing/controller"
	"github.com/krittawatcode/go-testing/service"
)

func main() {
	server := gin.Default()
	v1 := server.Group("/v1")
	{
		v1.POST("/eat/:commission", func(c *gin.Context) {
			var dutyService service.DutyService = service.SoldierDutyService(c)
			fmt.Printf("%+v \n", dutyService)
			fmt.Println("dutyService address is : ", &dutyService)
			var soldierController controller.SoldierController = controller.SoldierHandler(&dutyService)
			fmt.Printf("%+v \n", soldierController)
			soldierController.Eat(c)
		})
	}
	port := "8080"
	server.Run(":" + port)
}
