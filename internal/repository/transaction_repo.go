package repository

import (
	"context"
	"github.com/a-melchikov/avito-shop/internal/models"
)

type TransactionRepository struct {
	db DB
}

func NewTransactionRepository(db DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) AddTransaction(ctx context.Context, fromUserID int, toUserID int, amount int) error {
	query := `INSERT INTO coin_transactions (from_user_id, to_user_id, amount) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, fromUserID, toUserID, amount)
	return err
}

func (r *TransactionRepository) GetTransactionsByUser(ctx context.Context, userID int) ([]models.CoinTransaction, error) {
	query := `SELECT id, from_user_id, to_user_id, amount, transaction_date 
			  FROM coin_transactions 
			  WHERE from_user_id = $1 OR to_user_id = $1`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []models.CoinTransaction{}
	for rows.Next() {
		transaction := models.CoinTransaction{}
		if err := rows.Scan(&transaction.ID, &transaction.FromUserID, &transaction.ToUserID, &transaction.Amount, &transaction.TransactionDate); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
