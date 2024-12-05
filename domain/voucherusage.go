package domain

import (
	"project/util"
	"time"
)

type VoucherUsage struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	VoucherID uint      `gorm:"not null" json:"voucher_id"`
	UsedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"used_at"`
}

func VoucherUsageSeed() []VoucherUsage {
	return []VoucherUsage{
		{ID: 1, VoucherID: 1, UsedAt: util.Date("2024-11-16")},
		{ID: 2, VoucherID: 2, UsedAt: util.Date("2024-11-17")},
		{ID: 3, VoucherID: 3, UsedAt: util.Date("2024-11-18")},
		{ID: 4, VoucherID: 1, UsedAt: util.Date("2024-11-19")},
		{ID: 5, VoucherID: 2, UsedAt: util.Date("2024-11-20")},
	}
}
