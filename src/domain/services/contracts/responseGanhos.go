package contracts

type ResponseGanhos struct {
	ID                 int64   `json:"id"`
	Descricao          string  `json:"descricao"`
	Valor              float64 `json:"valor"`
	Tipo               string  `json:"tipo"`
	Data               string  `json:"data"`
	Categoria          string  `json:"categoria"`
	UsuarioIDCreatedAt int64   `json:"usuario_id_created_at"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
}
