package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"simple-service/config"
	"simple-service/storage"
)

const scope = "storage:postgresql_db"

type Store struct {
	db        *sqlx.DB
	log       *logrus.Entry
	cfg       *config.Config
	ProductR  *ProductRepo
	PriceR    *PriceRepo
	ShopR     *ShopRepo
	CategoryR *CategoryRepo
}

func (s *Store) Tx() (*sqlx.Tx, error) {
	return s.db.Beginx()
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

func NewStore(cfg *config.Config, log *logrus.Logger, db *sqlx.DB) (s *Store) {
	s = &Store{
		db:  db,
		log: log.WithFields(logrus.Fields{"scope": scope, "env": cfg.Environment}),
		cfg: cfg,
	}

	return s
}
