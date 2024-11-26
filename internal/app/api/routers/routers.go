package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/qvvan/dragonfit/internal/app/api/v1"
	"github.com/qvvan/dragonfit/internal/config"
)

func InitRouters(cfg *config.Config, v1Manager *v1.Manager) *gin.Engine {
	gin.SetMode(cfg.Debug)
	router := gin.Default()

	api := router.Group("/api")
	{
		RegisterPublicRoutes(api, v1Manager)
		RegisterPrivateRoutes(api, v1Manager)
	}

	return router
}
