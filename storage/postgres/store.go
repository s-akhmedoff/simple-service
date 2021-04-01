package postgres

import "simple-service/storage"

type Store struct {
	//db
	//log
	//cfg
	ProductR  *ProductRepo
	PriceR    *PriceRepo
	ShopR     *ShopRepo
	CategoryR *CategoryRepo
}

func (s *Store) Product() storage.ProductRepository {
	if s.ProductR == nil {
		s.ProductR = &ProductRepo{s: s}
	}

	return s.ProductR
}

func (s *Store) Price() storage.PriceRepository {
	if s.PriceR == nil {
		s.PriceR = &PriceRepo{s: s}
	}

	return s.PriceR
}

func (s *Store) Shop() storage.ShopRepository {
	if s.ShopR == nil {
		s.ShopR = &ShopRepo{s: s}
	}

	return s.ShopR
}

func (s *Store) Category() storage.CategoryRepository {
	if s.CategoryR == nil {
		s.CategoryR = &CategoryRepo{s: s}
	}

	return s.CategoryR
}

func NewStore() (s *Store) {
	s = &Store{}

	return
}
