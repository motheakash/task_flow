package models

import (
	"task_flow/pkg/base_models"
)

type UserModel struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Username string `json:"username" gorm:"type:varchar(30);not null;unique"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
	Role     string `json:"role" gorm:"type:varchar(20);not null"`
	base_models.BaseModel
}

func (UserModel) TableName() string {
	return "Users" // Use a custom table name (e.g., "users" instead of "user_models")
}
