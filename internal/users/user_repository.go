package repository

import (
    "context"
    "database/sql"
    "your-project/internal/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) GetUser(ctx context.Context, id uint) (*models.User, error) {
    return nil, nil
}
