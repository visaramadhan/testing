package repository

import (
	"project/domain"

	"gorm.io/gorm"
)

// VoucherUsageRepository adalah struct yang digunakan untuk operasi terkait VoucherUsage di database
type VoucherUsageRepository struct {
	db *gorm.DB
}

// NewVoucherUsageRepository mengembalikan instance baru dari VoucherUsageRepository
func NewVoucherUsageRepository(db *gorm.DB) *VoucherUsageRepository {
	return &VoucherUsageRepository{db: db}
}

// Save menyimpan VoucherUsage ke dalam database
func (r *VoucherUsageRepository) Save(voucherUsage *domain.VoucherUsage) error {
	return r.db.Create(voucherUsage).Error
}

// FindByID mencari VoucherUsage berdasarkan ID
func (r *VoucherUsageRepository) FindByID(id uint) (*domain.VoucherUsage, error) {
	var voucherUsage domain.VoucherUsage
	err := r.db.First(&voucherUsage, id).Error
	if err != nil {
		return nil, err
	}
	return &voucherUsage, nil
}

// FindAll mengambil semua VoucherUsage dari database
func (r *VoucherUsageRepository) FindAll() ([]domain.VoucherUsage, error) {
	var voucherUsages []domain.VoucherUsage
	err := r.db.Find(&voucherUsages).Error
	if err != nil {
		return nil, err
	}
	return voucherUsages, nil
}

// FindByUserID mencari VoucherUsage berdasarkan ID User (misalnya untuk mendapatkan semua penggunaan voucher oleh user tertentu)
func (r *VoucherUsageRepository) FindByUserID(userID uint) ([]domain.VoucherUsage, error) {
	var voucherUsages []domain.VoucherUsage
	err := r.db.Where("user_id = ?", userID).Find(&voucherUsages).Error
	if err != nil {
		return nil, err
	}
	return voucherUsages, nil
}

// FindByVoucherID mencari VoucherUsage berdasarkan ID Voucher (misalnya untuk mendapatkan semua penggunaan voucher tertentu)
func (r *VoucherUsageRepository) FindByVoucherID(voucherID uint) ([]domain.VoucherUsage, error) {
	var voucherUsages []domain.VoucherUsage
	err := r.db.Where("voucher_id = ?", voucherID).Find(&voucherUsages).Error
	if err != nil {
		return nil, err
	}
	return voucherUsages, nil
}
