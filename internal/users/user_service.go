package service

import (
    "context"
    "your-project/internal/models"
    "your-project/internal/repository"
)

type UserService struct {
    userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUser(ctx context.Context, id uint) (*models.User, error) {
    return s.userRepo.GetUser(ctx, id)
}
