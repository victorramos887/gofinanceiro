package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/gofinanceiro/src/domain/models"
	"github.com/victorramos887/gofinanceiro/src/domain/repository"
	"github.com/victorramos887/gofinanceiro/src/domain/services/contracts"
)

// CreateGanhoHandler godoc
// @Summary      Cria um novo ganho
// @Description  Cria um novo registro de ganho com os dados fornecidos no corpo da requisição
// @Tags         Ganhos
// @Accept       json
// @Produce      json
// @Param        ganho  body  contracts.RequestGanhos  true  "Dados do ganho a ser criado"
// @Success      201  {object}  map[string]interface{}  "Ganho criado com sucesso"
// @Failure      400  {object}  map[string]string       "Erro nos dados enviados"
// @Failure      500  {object}  map[string]string       "Erro interno no servidor"
// @Router       /api/v1/ganhos [post]
func CreateGanhoHandler(c *gin.Context) {
	h := repository.NewRepositoryGanhos(db)
	if h == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository not initialized"})
		return
	}
	var req contracts.RequestGanhos
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if req.Descricao == "" || req.Valor <= 0 || req.Tipo == "" || req.Data.Equal((time.Time{})) || req.Categoria == "" || req.UsuarioIDCreatedAt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	if req.UsuarioIDCreatedAt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ganho := models.Ganhos{
		Descricao: req.Descricao,
		Valor:     req.Valor,
		Tipo:      req.Tipo,
		Data:      req.Data,
	}

	createdGanho, err := h.CreateGanho(&ganho)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ganho"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Ganho created successfully", "data": createdGanho})
}
