package storage

import "simple-service/models"

type Storage interface {
	Product() ProductRepository
	Price() PriceRepository
	Category() CategoryRepository
	Shop() ShopRepository
}

type ProductRepository interface {
	Create(product *models.NewProduct) error
	Read() (*[]models.Product, error)
	ReadBySKI(SKI string) (*models.Product, error)
	ReadByType(productType string) (*models.Product, error)
	Update(ID string, product *models.Product) error
	Delete(ID string) error
}

type PriceRepository interface {
	Create(price *models.NewPrice) error
	Read() (*[]models.Price, error)
	ReadByProductID(ProductID string) (*models.Price, error)
	Update(ID string, price *models.NewPrice) error
	Delete(ID string) error
}

type CategoryRepository interface {
	Create(category *models.NewCategory) error
	Read() (*[]models.Category, error)
	ReadByName(name string) (*models.Category, error)
	Update(ID string, category *models.Category) error
	Delete(ID string) error
}

type ShopRepository interface {
	Create(shop *models.NewShop) error
	Read() (*[]models.Shop, error)
	ReadBy(name string) (*models.Shop, error)
	Update(ID string) error
	Delete(ID string) error
}
