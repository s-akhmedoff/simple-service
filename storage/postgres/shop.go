package postgres

import (
	"github.com/jmoiron/sqlx"
	"simple-service/models"
)

type ShopRepo struct {
	s *Store
}

func (s ShopRepo) Create(tx *sqlx.Tx, shop *models.NewShop) error {
	panic("implement me")
}

func (s ShopRepo) Read() (*[]models.Shop, error) {
	panic("implement me")
}

func (s ShopRepo) ReadByName(name string) (*models.Shop, error) {
	panic("implement me")
}

func (s ShopRepo) Update(tx *sqlx.Tx, ID string) error {
	panic("implement me")
}

func (s ShopRepo) Delete(tx *sqlx.Tx, ID string) error {
	panic("implement me")
}
