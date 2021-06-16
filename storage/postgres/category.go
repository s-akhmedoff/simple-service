package postgres

import (
	"github.com/jmoiron/sqlx"
	"simple-service/models"
)

type CategoryRepo struct {
	s *Store
}

func (c CategoryRepo) Create(tx *sqlx.Tx, category *models.NewCategory) error {
	panic("implement me")
}

func (c CategoryRepo) Read() (*[]models.Category, error) {
	panic("implement me")
}

func (c CategoryRepo) ReadByName(name string) (*models.Category, error) {
	panic("implement me")
}

func (c CategoryRepo) Update(tx *sqlx.Tx, ID string, category *models.Category) error {
	panic("implement me")
}

func (c CategoryRepo) Delete(tx *sqlx.Tx, ID string) error {
	panic("implement me")
}
