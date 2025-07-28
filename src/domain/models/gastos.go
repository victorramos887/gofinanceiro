package models

type Gastos struct {
	ID                 int64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Descricao          string  `json:"descricao" gorm:"type:varchar(255);not null"`
	Valor              float64 `json:"valor" gorm:"type:decimal(10,2);not null"`
	Data               string  `json:"data" gorm:"type:date;not null"`
	Categoria          string  `json:"categoria" gorm:"type:varchar(100);not null"`
	Parcelas           int     `json:"parcelas" gorm:"type:int;default:1"`
	DataInicio         string  `json:"data_inicio" gorm:"type:date;"`
	FormaPagamento     string  `json:"forma_pagamento" gorm:"type:varchar(50);"`
	Observacao         string  `json:"observacao" gorm:"type:text;null"`
	UsuarioIDCreatedAt int64   `json:"usuario_id" gorm:"type:bigint;not null"`
	UsuarioIDUpdatedAt int64   `json:"usuario_id_updated_at" gorm:"type:bigint;not null"`
	CreatedAt          string  `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt          string  `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt          string  `json:"deleted_at,omitempty" gorm:"type:timestamp;null"`
}

func (Gastos) TableName() string {
	return "gastos"
}
