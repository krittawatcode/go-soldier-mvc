package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-soldier-mvc/service"
)

type SoldierController interface {
	Eat(c *gin.Context)
}

type soldierController struct {
	dutyService service.DutyService
}

func SoldierHandler(dutyService *service.DutyService) SoldierController {
	return &soldierController{
		dutyService: *dutyService,
	}
}

func (s *soldierController) Eat(c *gin.Context) {
	commission, _ := strconv.Atoi(c.Param("commission"))
	s.dutyService.EatTax(c, commission)
}
