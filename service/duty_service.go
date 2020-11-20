package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-testing/cores"
)

type DutyService interface {
	EatTax(c *gin.Context, commission int)
}

type soldierInfo struct {
	Rank       string
	Wife       int
	Salary     int
	Home       bool
	Car        bool
	Corruption bool
}

func SoldierDutyService(c *gin.Context) DutyService {
	var soldier cores.Solider
	if err := c.ShouldBind(&soldier); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Printf("%+v \n", soldier)
	return &soldierInfo{
		Rank:       soldier.Rank,
		Wife:       soldier.Wife,
		Salary:     soldier.Salary,
		Home:       soldier.Home,
		Car:        soldier.Car,
		Corruption: soldier.Corruption,
	}
}

func (s *soldierInfo) EatTax(c *gin.Context, commission int) {
	fmt.Println("EatTax")
	fmt.Println(*s)

	if s.Rank != "elite" {
		s.Salary -= commission
	}
	if s.Corruption == true {
		s.getPromote("elite")
	}
	// jsonSoldier, err := json.Marshal(s)
	// if err != nil {
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"resp": s})
}

func (s *soldierInfo) getPromote(newrank string) {
	s.Rank = newrank
	s.Car = true
	s.Home = true
	s.Wife += 10
	s.Salary += s.Salary * 100
}
