package repository

import "github.com/victorramos887/gofinanceiro/src/domain/models"

func (r *Repository) GetGastosByID(id int64) (*models.Gastos, error) {
	var gastos models.Gastos
	if err := r.db.First(&gastos, id); err != nil {
		return nil, err
	}
	return &gastos, nil
}


func (r *Repository) CreateGasto(gasto *models.Gastos) error {
	if err := r.db.Create(gasto); err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateGasto(gasto *models.Gastos) error {
	if err := r.db.Save(gasto); err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteGasto(id int64) error {
	if err := r.db.Delete(&models.Gastos{}, id); err != nil {
		return err
	}
	return nil
}

func (r *Repository) ListGastos() ([]models.Gastos, error) {
	var gastos []models.Gastos
	if err := r.db.Find(&gastos); err != nil {
		return nil, err
	}
	return gastos, nil
}