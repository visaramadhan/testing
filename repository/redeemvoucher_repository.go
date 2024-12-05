package repository

import (
	"gorm.io/gorm"
	"project/domain"
)

type RedeemVoucherRepository struct {
	db *gorm.DB
}

func NewRedeemVoucherRepository(db *gorm.DB) *RedeemVoucherRepository {
	return &RedeemVoucherRepository{db}
}

func (r *RedeemVoucherRepository) Create(customer domain.Customer) error {
	return r.db.Create(&customer).Error
}
