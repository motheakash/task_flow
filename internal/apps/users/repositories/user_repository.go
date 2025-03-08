package repositories

import (
	"task_flow/internal/apps/users/models/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.UserModel) error {
	return r.DB.Create(user).Error
}
