package storage

import (
	"github.com/jmoiron/sqlx"
	"simple-service/models"
)

type Storage interface {
	Tx() (*sqlx.Tx, error)
	Product() ProductRepository
	Price() PriceRepository
	Category() CategoryRepository
	Shop() ShopRepository
}

type ProductRepository interface {
	Create(tx *sqlx.Tx, product *models.NewProduct) error
	Read() ([]*models.Product, error)
	ReadBySKI(SKI string) (*models.Product, error)
	ReadByType(productType string) (*models.Product, error)
	ReadByID(productID string) (*models.Product, error)
	Update(tx *sqlx.Tx, ID string, product *models.Product) error
	Delete(tx *sqlx.Tx, ID string) error
}

type PriceRepository interface {
	Create(tx *sqlx.Tx, price *models.NewPrice) error
	Read() (*[]models.Price, error)
	ReadByProductID(ProductID string) (*models.Price, error)
	Update(tx *sqlx.Tx, ID string, price *models.NewPrice) error
	Delete(tx *sqlx.Tx, ID string) error
}

type CategoryRepository interface {
	Create(tx *sqlx.Tx, category *models.NewCategory) error
	Read() (*[]models.Category, error)
	ReadByName(name string) (*models.Category, error)
	Update(tx *sqlx.Tx, ID string, category *models.Category) error
	Delete(tx *sqlx.Tx, ID string) error
}

type ShopRepository interface {
	Create(tx *sqlx.Tx, shop *models.NewShop) error
	Read() (*[]models.Shop, error)
	ReadByName(name string) (*models.Shop, error)
	Update(tx *sqlx.Tx, ID string) error
	Delete(tx *sqlx.Tx, ID string) error
}
