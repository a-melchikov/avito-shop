package repository

import (
	"context"
	"github.com/a-melchikov/avito-shop/internal/models"
)

type UserRepository struct {
	db DB
}

func NewUserRepository(db DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (username, password_hash, registration_date) 
			  VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(ctx, query, user.Username, user.PasswordHash, user.RegistrationDate).Scan(&user.ID)
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT id, username, password_hash, registration_date FROM users WHERE username = $1`
	user := &models.User{}
	err := r.db.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.RegistrationDate)
	if err != nil {
		return nil, err
	}
	return user, nil
}
