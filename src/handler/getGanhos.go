package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/gofinanceiro/src/domain/repository"
)


// GetGanhosHandler godoc
// @Summary      Retrieve a ganho by ID
// @Description  Fetches a ganho record based on the provided ID parameter
// @Tags         Ganhos
// @Accept       json
// @Produce      json
// @Param        id  path  int64  true  "ID of the ganho to retrieve"
// @Success      200  {object}  map[string]interface{}  "Ganho retrieved successfully"
// @Failure      400  {object}  map[string]string       "Invalid ID format"
// @Failure      404  {object}  map[string]string       "Ganho not found"
// @Failure      500  {object}  map[string]string       "Internal server error"
// @Router       /api/v1/ganhos/{id} [get]
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
