package router

import (
	"github.com/gin-gonic/gin"
	"github.com/victorramos887/gofinanceiro/src/domain/services/contracts"
	"github.com/victorramos887/gofinanceiro/src/handler"
)

func InitializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	v1 := router.Group(basePath)

	v1.GET("/ganhos", handler.CreateGanho(contracts.RequestGanhos{}))
}
