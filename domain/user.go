package domain

type User struct {
	ID            int            `gorm:"type varchar(255); primary key; not null; unique" json:"user_id" binding:"required"`
	Name          string         `gorm:"type varchar(255); not null" json:"name" binding:"required"`
	Email         string         `gorm:"type varchar(255); not null; unique" json:"email" binding:"required,email"`
	Password      string         `gorm:"type varchar(255); not null" json:"password" binding:"required"`
	Address       string         `gorm:"type varchar(255); not null" json:"address" binding:"required"`
	Voucher       []Voucher      `gorm:"many2many:user_vouchers" json:"vouchers"`
	VoucherUsages []VoucherUsage `gorm:"foreignKey:UserID" json:"voucher_usages"` // Relasi dengan VoucherUsage

}
