package contracts

import "time"

type RequestGastos struct {
	Descricao          string    `json:"descricao"`
	Valor              float64   `json:"valor"`
	Data               time.Time `json:"data"`
	Categoria          string    `json:"categoria"`
	Parcelas           int       `json:"parcelas"`
	DataInicio         time.Time `json:"data_inicio"`
	FormaPagamento     string    `json:"forma_pagamento"`
	Observacao         string    `json:"observacao"`
	UsuarioIDCreatedAt int64     `json:"usuario_id_created_at"`
	UsuarioIDUpdatedAt int64     `json:"usuario_id_updated_at"`
}
