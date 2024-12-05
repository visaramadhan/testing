package domain

type DiscountVoucher struct {
<<<<<<< HEAD
	ID        uint     `gorm:"primaryKey;autoIncrement" json:"-"`
	VoucherID uint     `json:"-"`
	Rate      *float32 `gorm:"precision:5;scale:2" json:"rate"`
	Amount    *uint32  `gorm:"precision:4" json:"amount"`
	Condition `gorm:"embedded" json:"terms_and_condition"`
=======
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	VoucherID uint
	Rate      float32 `gorm:"precision:5;scale:2" json:"rate"`
	Amount    uint32  `gorm:"precision:4" json:"amount"`
	Condition `gorm:"embedded" json:"-"`
>>>>>>> da0e613 (Initial commit)
}
