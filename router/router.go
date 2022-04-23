package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"metro/controller"
)

func SetRouters() *gin.Engine {
	r := gin.New()

	r.GET("/", func(context *gin.Context) {
		context.String(200, "metro go go go!!!")
	})

	//注册api文档
	r.GET("/swagger/*ang", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	v1.POST("/metros", controller.AddMetros)

	return r
}
