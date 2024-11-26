package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/qvvan/dragonfit/internal/app/api/v1"
)

func RegisterPrivateRoutes(group *gin.RouterGroup, v1 *v1.Manager) {
	// protected := group.Group("/protected")
	// // 	protected.Use(middleware.AuthMiddleware())
	// {
	// 	protected.GET("/profile")
	// 	protected.GET("/users")
	// }
}
