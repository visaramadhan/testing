package service

import (
	"project/domain"
	"project/repository"
)

type VoucherUsageServiceInterface interface {
	UseVoucher(voucherUsage *domain.VoucherUsage) error
	GetVoucherUsagesByUserID(userID uint) ([]domain.VoucherUsage, error)
	GetVoucherUsagesByVoucherID(voucherID uint) ([]domain.VoucherUsage, error)
}

type VoucherUsageService struct {
	voucherUsageRepo *repository.VoucherUsageRepository
}

func NewVoucherUsageService(voucherUsageRepo *repository.VoucherUsageRepository) *VoucherUsageService {
	return &VoucherUsageService{
		voucherUsageRepo: voucherUsageRepo,
	}
}

func (s *VoucherUsageService) UseVoucher(voucherUsage *domain.VoucherUsage) error {
	// Panggil repository untuk menyimpan penggunaan voucher
	err := s.voucherUsageRepo.Save(voucherUsage)
	if err != nil {
		return err
	}
	return nil
}

// Implementasi method GetVoucherUsagesByUserID
func (s *VoucherUsageService) GetVoucherUsagesByUserID(userID uint) ([]domain.VoucherUsage, error) {
	return s.voucherUsageRepo.FindByUserID(userID)
}

// Implementasi method GetVoucherUsagesByVoucherID
func (s *VoucherUsageService) GetVoucherUsagesByVoucherID(voucherID uint) ([]domain.VoucherUsage, error) {
	return s.voucherUsageRepo.FindByVoucherID(voucherID)
}
