package repository_test

import (
	"context"
	"github.com/a-melchikov/avito-shop/internal/repository"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestProductRepository_GetAllProducts(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := repository.NewProductRepository(mock)

	mock.ExpectQuery("SELECT id, name, price FROM products").
		WillReturnRows(mock.NewRows([]string{"id", "name", "price"}).
			AddRow(1, "Product 1", 100).
			AddRow(2, "Product 2", 200))

	products, err := repo.GetAllProducts(context.Background())
	require.NoError(t, err)
	require.Len(t, products, 2)
	require.Equal(t, "Product 1", products[0].Name)
	require.Equal(t, 200, products[1].Price)
	require.NoError(t, mock.ExpectationsWereMet())
}
