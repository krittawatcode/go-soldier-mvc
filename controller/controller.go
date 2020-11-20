package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-testing/service"
)

type SoldierController interface {
	Eat(c *gin.Context)
}

type soldierController struct {
	dutyService service.DutyService
}

func SoldierHandler(dutyService *service.DutyService) SoldierController {
	fmt.Println("Soldier Handler")
	fmt.Printf("%+v \n", *dutyService)
	return &soldierController{
		dutyService: *dutyService,
	}
}

func (s *soldierController) Eat(c *gin.Context) {
	fmt.Println("Eat func in controller :", s)
	commission, _ := strconv.Atoi(c.Param("commission"))
	s.dutyService.EatTax(c, commission)
}
