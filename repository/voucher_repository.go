package repository

import (
	"gorm.io/gorm"
<<<<<<< HEAD
	"log"
=======
>>>>>>> da0e613 (Initial commit)
	"project/domain"
)

type VoucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) *VoucherRepository {
	return &VoucherRepository{db}
}

func (r *VoucherRepository) Create(voucher domain.Voucher) error {
<<<<<<< HEAD
	log.Println(voucher)
	return nil
	//return r.db.Create(&voucher).Error
=======
	return r.db.Create(&voucher).Error
>>>>>>> da0e613 (Initial commit)
}

func (r *VoucherRepository) Get(id uint) (domain.Voucher, error) {
	var voucher domain.Voucher
	err := r.db.First(&voucher, id).Error
	return voucher, err
}

func (r *VoucherRepository) Update(voucher domain.Voucher) error {
	return r.db.Save(&voucher).Error
}

func (r *VoucherRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Voucher{}, id).Error
}
