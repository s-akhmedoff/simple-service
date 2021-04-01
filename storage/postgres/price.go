package postgres

import "simple-service/models"

type PriceRepo struct {
	s *Store
}

func (p PriceRepo) Create(price *models.NewPrice) error {
	panic("implement me")
}

func (p PriceRepo) Read() (*[]models.Price, error) {
	panic("implement me")
}

func (p PriceRepo) ReadByProductID(ProductID string) (*models.Price, error) {
	panic("implement me")
}

func (p PriceRepo) Update(ID string, price *models.NewPrice) error {
	panic("implement me")
}

func (p PriceRepo) Delete(ID string) error {
	panic("implement me")
}
