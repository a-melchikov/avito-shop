package repository

import (
	"context"
	"github.com/a-melchikov/avito-shop/internal/models"
)

type PurchaseRepository struct {
	db DB
}

func NewPurchaseRepository(db DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

func (r *PurchaseRepository) AddPurchase(ctx context.Context, userID, productID, quantity int) error {
	query := `INSERT INTO user_purchases (user_id, product_id, quantity) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, userID, productID, quantity)
	return err
}

func (r *PurchaseRepository) GetPurchasesByUser(ctx context.Context, userID int) ([]models.Purchase, error) {
	query := `SELECT id, user_id, product_id, quantity, purchase_date FROM user_purchases WHERE user_id = $1`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	purchases := []models.Purchase{}
	for rows.Next() {
		purchase := models.Purchase{}
		if err := rows.Scan(&purchase.ID, &purchase.UserID, &purchase.ProductID, &purchase.Quantity, &purchase.PurchaseDate); err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}
	return purchases, nil
}
