package repository

import (
	"context"
)

type WalletRepository struct {
	db DB
}

func NewWalletRepository(db DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) CreateWallet(ctx context.Context, userID int, initialBalance int) error {
	query := `INSERT INTO wallets (user_id, balance) VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, userID, initialBalance)
	return err
}

func (r *WalletRepository) GetBalance(ctx context.Context, userID int) (int, error) {
	query := `SELECT balance FROM wallets WHERE user_id = $1`
	var balance int
	err := r.db.QueryRow(ctx, query, userID).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, userID int, newBalance int) error {
	query := `UPDATE wallets SET balance = $1 WHERE user_id = $2`
	_, err := r.db.Exec(ctx, query, newBalance, userID)
	return err
}
