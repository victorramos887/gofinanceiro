package models

import (
	"time"

	"gorm.io/gorm"
)

type Ganhos struct {
	ID                 int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Descricao          string         `json:"descricao" gorm:"type:varchar(255)"`
	Valor              float64        `json:"valor" gorm:"type:decimal(10,2);not null"`
	Tipo               string         `json:"tipo" gorm:"type:varchar(50);not null"`
	Data               time.Time      `json:"data" gorm:"type:date;not null"`
	Categoria          string         `json:"categoria" gorm:"type:varchar(100);not null"`
	Observacao         string         `json:"observacao" gorm:"type:text"`
	UsuarioIDCreatedAt int64          `json:"usuario_id" gorm:"type:bigint;not null"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (Ganhos) TableName() string {
	return "ganhos"
}
