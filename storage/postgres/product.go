package postgres

import (
	"github.com/jmoiron/sqlx"
	"simple-service/models"
	"simple-service/utils"
	"time"
)

type ProductRepo struct {
	s *Store
}

func (p ProductRepo) Create(tx *sqlx.Tx, product *models.NewProduct) error {
	if product == nil {
		p.s.log.Error(errNilPointerReference)

		return errNilPointerReference
	}

	query := `INSERT INTO product(id,sku,name,type,uri,description,is_active) VALUES ($1,$2,$3,$4,$5,$6,$7)`

	_, err := tx.Exec(query,
		utils.GenUUID(),
		product.SKI,
		product.Name,
		product.ProductType,
		product.URI,
		product.Description,
		product.IsActive)
	if err != nil {
		p.s.log.Error(err)

		return err
	}

	return nil
}

func (p ProductRepo) Read() ([]*models.Product, error) {
	var products []*models.Product

	query := `SELECT * FROM product`

	rows, err := p.s.db.Queryx(query)
	if err != nil {
		p.s.log.Error(err)

		return nil, err
	}

	defer func() { _ = rows.Close() }()

	for rows.Next() {
		product := new(models.Product)

		err := rows.StructScan(product)
		if err != nil {
			p.s.log.Error(err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (p ProductRepo) ReadBySKI(SKI string) (*models.Product, error) {
	if SKI == "" {
		p.s.log.Error(errEmptyArgument)

		return nil, errEmptyArgument
	}

	product := new(models.Product)

	query := `SELECT * FROM product WHERE sku=?`

	err := p.s.db.QueryRowx(query, SKI).StructScan(product)
	if err != nil {
		p.s.log.Error(err)

		return nil, err
	}

	return product, nil
}

func (p ProductRepo) ReadByType(productType string) (*models.Product, error) {
	if productType == "" {
		p.s.log.Error(errEmptyArgument)

		return nil, errEmptyArgument
	}

	product := new(models.Product)

	query := `SELECT * FROM product WHERE type=?`

	err := p.s.db.QueryRowx(query, productType).StructScan(product)
	if err != nil {
		p.s.log.Error(err)

		return nil, err
	}

	return product, nil
}

func (p ProductRepo) Update(tx *sqlx.Tx, ID string, product *models.Product) error {
	if product == nil {
		p.s.log.Error(errNilPointerReference)

		return errNilPointerReference
	}

	if ID == "" {
		p.s.log.Error(errEmptyArgument)

		return errEmptyArgument
	}

	query := `UPDATE product SET 
                   sku=$2,
                   name=$3,
                   type=$4,
                   uri=$5,
                   description=$6,
                   is_active=$7,
                   updated_at=$8 
			WHERE id=$1`

	exec, err := tx.Exec(query,
		product.ID,
		product.SKI,
		product.Name,
		product.ProductType,
		product.URI,
		product.Description,
		product.IsActive,
		time.Now())
	if err != nil {
		p.s.log.Error(err)

		return err
	}

	if n, _ := exec.RowsAffected(); n == 0 {
		p.s.log.Error(errNoRowsAffected)

		return errNoRowsAffected
	}

	return nil
}

func (p ProductRepo) Delete(tx *sqlx.Tx, ID string) error {
	if ID == "" {
		p.s.log.Error(errEmptyArgument)

		return errEmptyArgument
	}

	query := `DELETE FROM product WHERE id=?`

	exec, err := tx.Exec(query, ID)
	if err != nil {
		p.s.log.Error(err)

		return err
	}

	if n, _ := exec.RowsAffected(); n == 0 {
		p.s.log.Error(errNoRowsAffected)

		return errNoRowsAffected
	}

	return nil
}
