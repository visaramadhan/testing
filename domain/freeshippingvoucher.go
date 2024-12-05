package domain

type FreeShippingVoucher struct {
<<<<<<< HEAD
	ID        uint `gorm:"primaryKey;autoIncrement" json:"-"`
	VoucherID uint `json:"-"`
	Condition `gorm:"embedded" json:"terms_and_condition"`
=======
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	VoucherID uint
	Condition `gorm:"embedded" json:"-"`
>>>>>>> da0e613 (Initial commit)
}
