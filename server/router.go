package server

import (
	"singo/api"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	// r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	// r.Use(middleware.Cors())
	// r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api")
	{
		v1.POST("ping", api.Ping)
		v1.POST("register", api.RegisterAPI)
		v1.POST("suspend", api.UpdateStudentAPI)
		v1.POST("retrievefornotifications", api.RetrieveForNotificationAPI)

		v1.GET("/commonstudets", api.RetrieveStudentsAPI)

	}
	return r
}
