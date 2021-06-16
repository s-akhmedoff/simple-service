package postgres

import (
	"github.com/jmoiron/sqlx"
	"simple-service/models"
)

type ProductRepo struct {
	s *Store
}

func (p ProductRepo) Create(tx *sqlx.Tx, product *models.NewProduct) error {
	panic("implement me")
}

func (p ProductRepo) Read() (*[]models.Product, error) {
	panic("implement me")
}

func (p ProductRepo) ReadBySKI(SKI string) (*models.Product, error) {
	panic("implement me")
}

func (p ProductRepo) ReadByType(productType string) (*models.Product, error) {
	panic("implement me")
}

func (p ProductRepo) Update(tx *sqlx.Tx, ID string, product *models.Product) error {
	panic("implement me")
}

func (p ProductRepo) Delete(tx *sqlx.Tx, ID string) error {
	panic("implement me")
}
