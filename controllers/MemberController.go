package controllers

import (
	"geo.org/gin-members/models"
	"geo.org/gin-members/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitMemberRoutes(router *gin.Engine) {
	router.GET("/members", getMembers)
	router.POST("/members", createMember)
	//router.GET("/members/:id", getMemberByID)
}

func getMembers(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetMembers())
}

func createMember(c *gin.Context) {
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, service.CreateMember(member))
}
