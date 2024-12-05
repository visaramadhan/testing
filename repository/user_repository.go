package repository

import (
	"gorm.io/gorm"
	"project/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Anda dapat menambahkan fungsi-fungsi lain sesuai kebutuhan
