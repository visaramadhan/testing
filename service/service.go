<<<<<<< HEAD
=======
package service

import "project/repository"

type Service struct {
	Voucher       VoucherService
	RedeemVoucher RedeemVoucherService
	VoucherUsage  *VoucherUsageService
}

func NewService(repo repository.Repository) Service {
	return Service{
		Voucher:       NewVoucherService(repo.Voucher),
		RedeemVoucher: NewRedeemVoucherService(repo.RedeemVoucher),
		VoucherUsage:  NewVoucherUsageService(&repo.VoucherUsage),
	}
}
>>>>>>> e69a3da (Voucher usages)
