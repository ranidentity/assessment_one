package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

func RegisterAPI(c *gin.Context) {
	var service service.StudentRegistrationService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func RetrieveStudentsAPI(c *gin.Context) {
	var service service.RetrieveStudentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.RetrieveStudent()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UpdateStudentAPI(c *gin.Context) {
	var service service.UpdateStudentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateStudent()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func RetrieveForNotificationAPI(c *gin.Context) {
	var service service.SendNotificationService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CheckNotificationTarget()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
