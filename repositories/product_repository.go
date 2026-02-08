package repositories

import (
	"app-kasir/models"
	"database/sql"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll(nameFilter string) ([]models.ProductResponse, error) {
	query := `
		SELECT 
			p.id, p.name, p.price, p.stock,
			c.id, c.name, c.description
		FROM products p
		JOIN categories c ON c.id = p.category_id
	`

	args := []interface{}{}
	if nameFilter != "" {
		query += " WHERE p.name ILIKE $1"
		args = append(args, "%"+nameFilter+"%")
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []models.ProductResponse{}

	for rows.Next() {
		var p models.ProductResponse
		p.Category = models.Category{}

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.Stock,
			&p.Category.ID,
			&p.Category.Name,
			&p.Category.Description,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepository) GetByID(id int) (*models.ProductResponse, error) {
	query := `
		SELECT 
			p.id, p.name, p.price, p.stock,
			c.id, c.name, c.description
		FROM products p
		JOIN categories c ON c.id = p.category_id
		WHERE p.id = $1
	`

	var p models.ProductResponse
	p.Category = models.Category{}

	err := repo.db.QueryRow(query, id).
		Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.Stock,
			&p.Category.ID,
			&p.Category.Name,
			&p.Category.Description,
		)

	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	return &p, err
}

func (repo *ProductRepository) Create(product *models.Product) error {
	query := `
        INSERT INTO products (name, price, stock, category_id)
        VALUES ($1, $2, $3, $4) RETURNING id
    `
	return repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).
		Scan(&product.ID)
}

func (repo *ProductRepository) Update(product *models.Product) error {
	query := `
        UPDATE products SET name=$1, price=$2, stock=$3, category_id=$4 WHERE id=$5
    `
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryID, product.ID)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	result, err := repo.db.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("product not found")
	}
	return nil
}
