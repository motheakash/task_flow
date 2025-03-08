package services

import (
	"task_flow/internal/apps/users/models"
	"task_flow/internal/apps/users/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *models.UserModel) error {
	return s.Repo.CreateUser(user)
}
