package service

import (
	"app/adapter/repository"

	"gorm.io/gorm"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(db),
	}
}
