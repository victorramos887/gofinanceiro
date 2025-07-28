package repository

import "github.com/victorramos887/gofinanceiro/src/domain/models"


type Repository struct {
	db interface {
		First(dest interface{}, conds ...interface{}) error
		Create(value interface{}) error
		Save(value interface{}) error
		Delete(value interface{}, conds ...interface{}) error
		Find(dest interface{}, conds ...interface{}) error
	}
}

func (r *Repository) GetGanhosByID(id int64) (*models.Ganhos, error) {
	var ganho models.Ganhos
	if err := r.db.First(&ganho, id); err != nil {
		return nil, err
	}
	return &ganho, nil
}

func (r *Repository) CreateGanho(ganho *models.Ganhos) error {
	if err := r.db.Create(ganho); err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateGanho(ganho *models.Ganhos) error {
	if err := r.db.Save(ganho); err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteGanho(id int64) error {
	if err := r.db.Delete(&models.Ganhos{}, id); err != nil {
		return err
	}
	return nil
}

func (r *Repository) ListGanhos() ([]models.Ganhos, error) {
	var ganhos []models.Ganhos
	if err := r.db.Find(&ganhos); err != nil {
		return nil, err
	}
	return ganhos, nil
}
