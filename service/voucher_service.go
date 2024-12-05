package service

import (
	"project/domain"
	"project/repository"
)

type VoucherService interface {
	Create(voucher domain.Voucher) error
	Get(id uint) (domain.Voucher, error)
	Update(voucher domain.Voucher) error
	Delete(id uint) error
}

type voucherService struct {
	repo repository.VoucherRepository
}

func NewVoucherService(repo repository.VoucherRepository) VoucherService {
	return &voucherService{repo}
}

func (s *voucherService) Create(voucher domain.Voucher) error {
	return s.repo.Create(voucher)
}

func (s *voucherService) Get(id uint) (domain.Voucher, error) {
	return s.repo.Get(id)
}

func (s *voucherService) Update(voucher domain.Voucher) error {
	return s.repo.Update(voucher)
}

func (s *voucherService) Delete(id uint) error {
	return s.repo.Delete(id)
}
