package contracts

import "time"

type RequestGanhos struct {
	Descricao          string    `json:"descricao"`
	Valor              float64   `json:"valor"`
	Tipo               string    `json:"tipo"`
	Data               time.Time `json:"data"`
	Categoria          string    `json:"categoria"`
	Observacao         string    `json:"observacao"`
	UsuarioIDCreatedAt int64     `json:"usuario_id"`
}
