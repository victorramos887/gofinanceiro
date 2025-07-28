package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/gofinanceiro/src/domain/repository"
)


// GetGastosByIDHandler godoc
// @Summary      Retrieve a gasto by ID
// @Description  Fetches a gasto record based on the provided ID parameter
// @Tags         Gastos
// @Accept       json
// @Produce      json
// @Param        id  path  int64  true  "ID of the gasto to retrieve"
// @Success      200  {object}  map[string]interface{}  "Gasto retrieved successfully"
// @Failure      400  {object}  map[string]string       "Invalid ID format"
// @Failure      404  {object}  map[string]string       "Gasto not found"
// @Failure      500  {object}  map[string]string       "Internal server error"
// @Router       /gastos/{id} [get]
func GetGastosByIDHandler(c *gin.Context) {
	h := repository.NewRepositoryGastos(db)
	if h == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository not initialized"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a valid integer"})
		return
	}

	gastos, err := h.GetGastosByID(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve gastos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gastos})
}
