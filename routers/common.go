package routers

import (
	"doki.life/controllers"
	"doki.life/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitCommonRouter(router *gin.RouterGroup) {
	router.GET("/ping", controllers.Ping)
}

func InitSwaggerRouter(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
