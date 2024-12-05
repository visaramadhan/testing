package service

import (
	"project/domain"
	"project/repository"
)

type RedeemVoucherService interface {
	Create(customer domain.Customer) error
}

type redeemVoucherService struct {
	repo repository.RedeemVoucherRepository
}

func NewRedeemVoucherService(repo repository.RedeemVoucherRepository) RedeemVoucherService {
	return &redeemVoucherService{repo}
}

func (s *redeemVoucherService) Create(customer domain.Customer) error {
	return s.repo.Create(customer)
}
