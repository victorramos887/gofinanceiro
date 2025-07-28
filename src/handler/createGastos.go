package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorramos887/gofinanceiro/src/domain/models"
	"github.com/victorramos887/gofinanceiro/src/domain/repository"
	"github.com/victorramos887/gofinanceiro/src/domain/services/contracts"
)


// CreateGastosHandler godoc
// @Summary      Cria um novo gasto
// @Description  Cria um novo registro de gastos com os dados fornecidos no corpo da requisição
// @Tags         Gastos
// @Accept       json
// @Produce      json
// @Param        gasto  body  contracts.RequestGastos  true  "Dados do gasto a ser criado"
// @Success      201  {object}  map[string]interface{}  "Gasto criado com sucesso"
// @Failure      400  {object}  map[string]string       "Erro nos dados enviados"
// @Failure      500  {object}  map[string]string       "Erro interno no servidor"
// @Router       /api/v1/gastos [post]
func CreateGastosHandler(c *gin.Context) {

	h := repository.NewRepositoryGastos(db)
	if h == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository not initialized"})
		return
	}

	// criar request de gasto var req contracts.RequestGastos
	var req contracts.RequestGastos
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if req.Descricao == "" || req.Valor <= 0 || req.Categoria == "" || req.FormaPagamento == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields: descricao, valor, categoria, forma_pagamento"})
		return
	}

	if req.UsuarioIDCreatedAt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	gasto := models.Gastos{
		Descricao:          req.Descricao,
		Valor:              req.Valor,
		Data:               req.Data,
		Categoria:          req.Categoria,
		Parcelas:           req.Parcelas,
		DataInicio:         req.DataInicio,
		FormaPagamento:     req.FormaPagamento,
		Observacao:         req.Observacao,
		UsuarioIDCreatedAt: req.UsuarioIDCreatedAt,
	}

	createGastos, err := h.CreateGastos(&gasto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create gasto"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Gasto created successfully", "data": createGastos})
}
