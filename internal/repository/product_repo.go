package repository

import (
	"context"
	"github.com/a-melchikov/avito-shop/internal/models"
)

type ProductRepository struct {
	db DB
}

func NewProductRepository(db DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	query := `SELECT id, name, price FROM products`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
