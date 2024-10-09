package router

import (
	"github.com/PandaPy/pginer/template/api/health"
	"github.com/PandaPy/pginer/template/initialize/config"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置主路由
func SetupRoutes(r *gin.Engine) {

	// 公共路由
	PublicGroup := r.Group(config.AppConfig.ROUTER_PREFIX)
	{
		health.RegisterRoutes(PublicGroup)
	}
	// 私有路由
	// PrivateGroup := r.Group(config.AppConfig.ROUTER_PREFIX)
	// {
	// 	// user.RegisterRoutes(PrivateGroup)
	// }
}
