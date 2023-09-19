package router

import "github.com/gin-gonic/gin"

// InitRouter
//
//	@Description:路由初始化
//	@return *gin.Engine
func InitRouter() *gin.Engine {
	routers := gin.Default()
	adminRouter := routers.Group("/", nil)
	{
		adminRouter.GET("/")
	}
	return routers
}
