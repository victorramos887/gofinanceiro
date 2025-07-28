package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/victorramos887/gofinanceiro/src/docs"
	"github.com/victorramos887/gofinanceiro/src/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group(basePath)

	v1.POST("/ganhos", handler.CreateGanhoHandler)
	v1.GET("/ganhos/:id", handler.GetGanhosHandler)
	v1.GET("/ganhos", handler.ListGanhosHandler)

	v1.POST("/gastos", handler.CreateGastosHandler)
	v1.GET("/gastos/:id", handler.GetGastosByIDHandler)

}
