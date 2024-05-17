package repository

import (
	"app/adapter/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CRUD
func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}
