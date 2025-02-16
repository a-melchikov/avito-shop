package repository_test

import (
	"context"
	"github.com/a-melchikov/avito-shop/internal/repository"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestPurchaseRepository_AddPurchase(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := repository.NewPurchaseRepository(mock)

	mock.ExpectExec("INSERT INTO user_purchases").
		WithArgs(1, 2, 3).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err = repo.AddPurchase(context.Background(), 1, 2, 3)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestPurchaseRepository_GetPurchasesByUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := repository.NewPurchaseRepository(mock)

	purchaseDate1, _ := time.Parse("2006-01-02", "2023-01-01")
	purchaseDate2, _ := time.Parse("2006-01-02", "2023-01-02")

	mock.ExpectQuery("SELECT id, user_id, product_id, quantity, purchase_date FROM user_purchases").
		WithArgs(1).
		WillReturnRows(mock.NewRows([]string{"id", "user_id", "product_id", "quantity", "purchase_date"}).
			AddRow(1, 1, 2, 3, purchaseDate1).
			AddRow(2, 1, 3, 1, purchaseDate2))

	purchases, err := repo.GetPurchasesByUser(context.Background(), 1)
	require.NoError(t, err)

	require.Len(t, purchases, 2)
	require.Equal(t, 3, purchases[0].Quantity)
	require.Equal(t, purchaseDate1, purchases[0].PurchaseDate)
	require.Equal(t, purchaseDate2, purchases[1].PurchaseDate)

	require.NoError(t, mock.ExpectationsWereMet())
}
