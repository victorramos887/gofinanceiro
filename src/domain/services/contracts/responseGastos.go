package contracts

type ResponseGastos struct {
	ID                 int64   `json:"id"`
	Descricao          string  `json:"descricao"`
	Valor              float64 `json:"valor"`
	Data               string  `json:"data"`
	Categoria          string  `json:"categoria"`
	Parcelas           int     `json:"parcelas"`
	DataInicio         string  `json:"data_inicio"`
	FormaPagamento     string  `json:"forma_pagamento"`
	Observacao         string  `json:"observacao"`
	UsuarioIDCreatedAt int64   `json:"usuario_id_created_at"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
}