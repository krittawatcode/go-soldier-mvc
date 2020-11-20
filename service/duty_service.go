package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-soldier-mvc/cores"
)

// DutyService for define all service
type DutyService interface {
	EatTax(c *gin.Context, commission int)
}

type soldierInfo struct {
	Rank       string `json:"rank"`
	Wife       int    `json:"wife"`
	Salary     int    `json:"salary"`
	Home       bool   `json:"home"`
	Car        bool   `json:"car"`
	Corruption bool   `json:"corruption"`
}

// SoldierDutyService work like a contructer in java
func SoldierDutyService(c *gin.Context) DutyService {
	var soldier cores.Solider
	if err := c.ShouldBind(&soldier); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
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
	if s.Rank != "General" { // Field Marshall ??
		s.Salary -= commission
	}
	if s.Corruption == true {
		s.getPromote("Elite")
	}
	c.JSON(http.StatusOK, gin.H{"resp": s})
}

func (s *soldierInfo) getPromote(newrank string) {
	s.Rank = newrank
	s.Car = true
	s.Home = true
	s.Wife += 10
	s.Salary += s.Salary * 100
}
