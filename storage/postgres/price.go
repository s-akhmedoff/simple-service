package postgres

import (
	"github.com/jmoiron/sqlx"
	"simple-service/models"
)

type PriceRepo struct {
	s *Store
}

func (p PriceRepo) Create(tx *sqlx.Tx, price *models.NewPrice) error {
	panic("implement me")
}

func (p PriceRepo) Read() (*[]models.Price, error) {
	panic("implement me")
}

func (p PriceRepo) ReadByProductID(ProductID string) (*models.Price, error) {
	panic("implement me")
}

func (p PriceRepo) Update(tx *sqlx.Tx, ID string, price *models.NewPrice) error {
	panic("implement me")
}

func (p PriceRepo) Delete(tx *sqlx.Tx, ID string) error {
	panic("implement me")
}
