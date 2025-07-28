package router

import "github.com/gin-gonic/gin"

func InitializeRoutes() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	initializeRoutes(r)
	r.Run(":8085")
}
