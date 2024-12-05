package domain

type Customer struct {
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string         `json:"name"`
	RewardPoints uint           `json:"reward_points"`
	Redemptions  []Voucher      `gorm:"many2many:redemptions;" json:"redemptions"`
	Vouchers     []Voucher      `gorm:"many2many:orders;" json:"vouchers"`
	VoucherUsage []VoucherUsage `json:"voucher_usage" gorm:"foreignKey:CustomerID"`
}

func CustomerSeed() []Customer {
	return []Customer{
		{Name: "Customer Satu", RewardPoints: 10},
		{Name: "Customer Dua", RewardPoints: 100},
		{Name: "Customer Tiga", RewardPoints: 50},
		{Name: "Customer Empat", RewardPoints: 30},
		{Name: "Customer Lima", RewardPoints: 25},
		{Name: "Customer Enam", RewardPoints: 100},
		{Name: "Customer Tujuh", RewardPoints: 500},
		{Name: "Customer Delapan", RewardPoints: 25},
		{Name: "Customer Sembilan", RewardPoints: 75},
		{Name: "Customer Sepuluh", RewardPoints: 200},
	}
}

func CustomerVoucherUsageSeed() []VoucherUsage {
	return []VoucherUsage{
		{ID: 1, VoucherID: 1},
		{ID: 2, VoucherID: 2},
		{ID: 3, VoucherID: 3},
		{ID: 4, VoucherID: 4},
		{ID: 5, VoucherID: 5},
		{ID: 6, VoucherID: 1},
		{ID: 7, VoucherID: 2},
		{ID: 8, VoucherID: 3},
		{ID: 9, VoucherID: 4},
		{ID: 10, VoucherID: 5},
	}
}
