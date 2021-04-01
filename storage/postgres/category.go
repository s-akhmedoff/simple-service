package postgres

import (
	"simple-service/models"
)

type CategoryRepo struct {
	s *Store
}

func (c CategoryRepo) Create(category *models.NewCategory) error {
	panic("implement me")
}

func (c CategoryRepo) Read() (*[]models.Category, error) {
	panic("implement me")
}

func (c CategoryRepo) ReadByName(name string) (*models.Category, error) {
	panic("implement me")
}

func (c CategoryRepo) Update(ID string, category *models.Category) error {
	panic("implement me")
}

func (c CategoryRepo) Delete(ID string) error {
	panic("implement me")
}
