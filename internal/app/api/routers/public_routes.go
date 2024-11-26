package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/qvvan/dragonfit/internal/app/api/v1"
)

func RegisterPublicRoutes(apiGroup *gin.RouterGroup, v1Manager *v1.Manager) {
	v1Group := apiGroup.Group("/v1")
	{
		v1Group.POST("/login", v1Manager.Login)
		v1Group.POST("/register", v1Manager.Register)
	}
}
