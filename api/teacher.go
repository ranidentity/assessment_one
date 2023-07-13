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
