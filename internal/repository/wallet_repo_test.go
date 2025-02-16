package repository_test

import (
	"context"
	"github.com/a-melchikov/avito-shop/internal/repository"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestWalletRepository_CreateWallet(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := repository.NewWalletRepository(mock)

	mock.ExpectExec("INSERT INTO wallets").
		WithArgs(1, 100).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = repo.CreateWallet(context.Background(), 1, 100)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestWalletRepository_GetBalance(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := repository.NewWalletRepository(mock)

	mock.ExpectQuery("SELECT balance FROM wallets").
		WithArgs(1).
		WillReturnRows(mock.NewRows([]string{"balance"}).AddRow(100))

	balance, err := repo.GetBalance(context.Background(), 1)
	require.NoError(t, err)
	require.Equal(t, 100, balance)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestWalletRepository_UpdateBalance(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := repository.NewWalletRepository(mock)

	mock.ExpectExec("UPDATE wallets").
		WithArgs(150, 1).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = repo.UpdateBalance(context.Background(), 1, 150)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
