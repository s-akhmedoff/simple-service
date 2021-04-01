package postgres

import (
	"simple-service/models"
)

type ProductRepo struct {
	s *Store
}

func (p ProductRepo) Create(product *models.NewProduct) error {
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

func (p ProductRepo) Update(ID string, product *models.Product) error {
	panic("implement me")
}

func (p ProductRepo) Delete(ID string) error {
	panic("implement me")
}
