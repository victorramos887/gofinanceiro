package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/gofinanceiro/src/domain/services/contracts"
)

func CreateGanho(service contracts.RequestGanhos) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"message": "Ganho created successfully", "data": service})
	}
}
