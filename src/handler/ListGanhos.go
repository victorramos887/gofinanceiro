package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/gofinanceiro/src/domain/repository"
	"github.com/victorramos887/gofinanceiro/src/domain/services/contracts"
)

// ListGanhosHandler godoc
// @Summary      Retrieve a ganho list
// @Description  Fetches a ganho record based List
// @Tags         Ganhos
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "Ganho retrieved successfully"
// @Failure      400  {object}  map[string]string       "Invalid ID format"
// @Failure      404  {object}  map[string]string       "Ganho not found"
// @Failure      500  {object}  map[string]string       "Internal server error"
// @Router       /ganhos/ [get]
func ListGanhosHandler(c *gin.Context) {
	h := repository.NewRepositoryGanhos(db)
	if h == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository not initialized"})
		return
	}

	ganhos, err := h.ListGanhos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ganhos"})
		return
	}

	var responseGanhos []contracts.ResponseGanhos
	for _, ganho := range ganhos {
		responseGanho := contracts.ResponseGanhos{
			ID:        ganho.ID,
			Descricao: ganho.Descricao,
			Valor:     ganho.Valor,
			Tipo:      ganho.Tipo,
			Data:      ganho.Data.Format("2006-01-02"),
			Categoria: ganho.Categoria,
			CreatedAt: ganho.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: ganho.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		responseGanhos = append(responseGanhos, responseGanho)
	}

	c.JSON(http.StatusOK, gin.H{"data": responseGanhos})
}
