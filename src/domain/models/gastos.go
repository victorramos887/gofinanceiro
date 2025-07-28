package models

import (
	"time"

	"gorm.io/gorm"
)

type Gastos struct {
	ID                 int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Descricao          string         `json:"descricao" gorm:"type:varchar(255);"`
	Valor              float64        `json:"valor" gorm:"type:decimal(10,2);"`
	Data               time.Time      `json:"data" gorm:"type:date;"`
	Categoria          string         `json:"categoria" gorm:"type:varchar(100);not null"`
	Parcelas           int            `json:"parcelas" gorm:"type:int;default:1"`
	DataInicio         time.Time      `json:"data_inicio" gorm:"type:date;"`
	FormaPagamento     string         `json:"forma_pagamento" gorm:"type:varchar(50);"`
	Observacao         string         `json:"observacao" gorm:"type:text;null"`
	UsuarioIDCreatedAt int64          `json:"usuario_id" gorm:"type:bigint;not null"`
	UsuarioIDUpdatedAt int64          `json:"usuario_id_updated_at" gorm:"type:bigint;not null"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (Gastos) TableName() string {
	return "gastos"
}
