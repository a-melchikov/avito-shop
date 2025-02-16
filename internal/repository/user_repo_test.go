package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/a-melchikov/avito-shop/internal/models"
	"github.com/a-melchikov/avito-shop/internal/repository"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_CreateUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := repository.NewUserRepository(mock)

	registrationDate, _ := time.Parse("2006-01-02", "2023-01-01")

	user := &models.User{
		Username:         "testuser",
		PasswordHash:     "hash",
		RegistrationDate: registrationDate,
	}

	mock.ExpectQuery("INSERT INTO users").
		WithArgs("testuser", "hash", registrationDate).
		WillReturnRows(mock.NewRows([]string{"id"}).AddRow(1))

	err = repo.CreateUser(context.Background(), user)
	require.NoError(t, err)
	require.Equal(t, 1, user.ID)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetUserByUsername(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := repository.NewUserRepository(mock)

	registrationDate, _ := time.Parse("2006-01-02", "2023-01-01")

	mock.ExpectQuery("SELECT id, username, password_hash, registration_date FROM users").
		WithArgs("testuser").
		WillReturnRows(mock.NewRows([]string{"id", "username", "password_hash", "registration_date"}).
			AddRow(1, "testuser", "hash", registrationDate))

	user, err := repo.GetUserByUsername(context.Background(), "testuser")
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "testuser", user.Username)
	require.Equal(t, registrationDate, user.RegistrationDate)
	require.NoError(t, mock.ExpectationsWereMet())
}
