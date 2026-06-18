package store

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/zagvozdeen/donely/internal/store/models"
)

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	err := s.pool.QueryRow(ctx, "SELECT * FROM users WHERE email = $1", email).Scan(
		&user.ID,
		&user.UUID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByUUID(ctx context.Context, uuid uuid.UUID) (*models.User, error) {
	user := &models.User{}
	err := s.pool.QueryRow(ctx, "SELECT * FROM users WHERE uuid = $1", uuid).Scan(
		&user.ID,
		&user.UUID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s *Store) CreateUser(ctx context.Context, user *models.User) error {
	return s.pool.QueryRow(
		ctx,
		"INSERT INTO users (uuid, first_name, last_name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING (id)",
		user.UUID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
}
