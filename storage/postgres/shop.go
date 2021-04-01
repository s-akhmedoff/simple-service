package postgres

import "simple-service/models"

type ShopRepo struct {
	s *Store
}

func (s ShopRepo) Create(shop *models.NewShop) error {
	panic("implement me")
}

func (s ShopRepo) Read() (*[]models.Shop, error) {
	panic("implement me")
}

func (s ShopRepo) ReadBy(name string) (*models.Shop, error) {
	panic("implement me")
}

func (s ShopRepo) Update(ID string) error {
	panic("implement me")
}

func (s ShopRepo) Delete(ID string) error {
	panic("implement me")
}
