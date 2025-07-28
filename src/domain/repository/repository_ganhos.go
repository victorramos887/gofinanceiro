package repository

import (
	"github.com/victorramos887/gofinanceiro/src/domain/models"
	"gorm.io/gorm"
)

type RepositoryGanhos struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryGanhos {
	return &RepositoryGanhos{db: db}
}

func (r *RepositoryGanhos) GetGanhosByID(id int64) (*models.Ganhos, error) {
	var ganho models.Ganhos
	if err := r.db.First(&ganho, id).Error; err != nil {
		return nil, err
	}
	return &ganho, nil
}

func (r *RepositoryGanhos) CreateGanho(ganho *models.Ganhos) error {
	if err := r.db.Create(ganho).Error; err != nil {
		return err
	}
	return nil
}

func (r *RepositoryGanhos) UpdateGanho(ganho *models.Ganhos) error {
	if err := r.db.Save(ganho).Error; err != nil {
		return err
	}
	return nil
}

func (r *RepositoryGanhos) DeleteGanho(id int64) error {
	if err := r.db.Delete(&models.Ganhos{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *RepositoryGanhos) ListGanhos() ([]models.Ganhos, error) {
	var ganhos []models.Ganhos
	if err := r.db.Find(&ganhos).Error; err != nil {
		return nil, err
	}
	return ganhos, nil
}
