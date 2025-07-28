package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/gofinanceiro/src/domain/repository"
)

func GetGanhosHandler(c *gin.Context) {
	h := repository.NewRepositoryGanhos(db)
	if h == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository not initialized"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}

	// Convert id from string to int64
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a valid integer"})
		return
	}

	ganhos, err := h.GetGanhosByID(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ganhos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ganhos})
}
