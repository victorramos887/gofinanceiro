package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/gofinanceiro/src/domain/repository"
	"github.com/victorramos887/gofinanceiro/src/domain/services/contracts"
)

// ListGastosHandler godoc
// @Summary      Retrieve a gasto list
// @Description  Fetches a gasto record based List
// @Tags         Gastos
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "Gasto retrieved successfully"
// @Failure      400  {object}  map[string]string       "Invalid ID format"
// @Failure      404  {object}  map[string]string       "Gasto not found"
// @Failure      500  {object}  map[string]string       "Internal server error"
// @Router       /gastos/ [get]
func ListGastosHandler(c *gin.Context) {
	h := repository.NewRepositoryGastos(db)
	if h == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository not initialized"})
		return
	}

	gastos, err := h.ListGastos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve gastos"})
		return
	}

	var responseGastos []contracts.ResponseGastos
	for _, gasto := range gastos {
		responseGasto := contracts.ResponseGastos{
			ID:                 gasto.ID,
			Descricao:          gasto.Descricao,
			Valor:              gasto.Valor,
			Data:               gasto.Data.Format("2006-01-02"),
			Categoria:          gasto.Categoria,
			Parcelas:           gasto.Parcelas,
			DataInicio:         gasto.DataInicio.Format("2006-01-02"),
			FormaPagamento:     gasto.FormaPagamento,
			Observacao:         gasto.Observacao,
			UsuarioIDCreatedAt: gasto.UsuarioIDCreatedAt,
			CreatedAt:          gasto.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:          gasto.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		responseGastos = append(responseGastos, responseGasto)
	}

	c.JSON(http.StatusOK, gin.H{"data": responseGastos})
}
