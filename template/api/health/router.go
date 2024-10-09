package health

import "github.com/gin-gonic/gin"

// RegisterRoutes 注册 User 模块的路由
func RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/health")
	{
		router.GET("", CheckHealth)
	}
}
