package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/a-melchikov/avito-shop/internal/repository"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestTransactionRepository_AddTransaction(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := repository.NewTransactionRepository(mock)

	fromUserID := 1
	toUserID := 2
	amount := 100

	mock.ExpectExec("INSERT INTO coin_transactions").
		WithArgs(fromUserID, toUserID, amount).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = repo.AddTransaction(context.Background(), fromUserID, toUserID, amount)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestTransactionRepository_GetTransactionsByUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := repository.NewTransactionRepository(mock)

	date1, _ := time.Parse("2006-01-02", "2023-01-01")
	date2, _ := time.Parse("2006-01-02", "2023-01-02")

	mock.ExpectQuery("SELECT id, from_user_id, to_user_id, amount, transaction_date FROM coin_transactions").
		WithArgs(1).
		WillReturnRows(mock.NewRows([]string{"id", "from_user_id", "to_user_id", "amount", "transaction_date"}).
			AddRow(1, 1, 2, 100, date1).
			AddRow(2, 3, 1, 50, date2))

	transactions, err := repo.GetTransactionsByUser(context.Background(), 1)
	require.NoError(t, err)
	require.Len(t, transactions, 2)
	require.Equal(t, 100, transactions[0].Amount)
	require.Equal(t, 50, transactions[1].Amount)
	require.NoError(t, mock.ExpectationsWereMet())
}
