package models

type Ganhos struct {
	ID                 int64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Descricao          string  `json:"descricao" gorm:"type:varchar(255);"`
	Valor              float64 `json:"valor" gorm:"type:decimal(10,2);not null"`
	Tipo               string  `json:"tipo" gorm:"type:varchar(50);not null"`
	Data               string  `json:"data" gorm:"type:date;not null"`
	Categoria          string  `json:"categoria" gorm:"type:varchar(100);not null"`
	Observacao         string  `json:"observacao" gorm:"type:text;null"`
	UsuarioIDCreatedAt int64   `json:"usuario_id" gorm:"type:bigint;not null"`
	CreatedAt          string  `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt          string  `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt          string  `json:"deleted_at,omitempty" gorm:"type:timestamp;null"`
}

func (Ganhos) TableName() string {
	return "ganhos"
}
