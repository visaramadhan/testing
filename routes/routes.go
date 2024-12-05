package routes

import (
	"project/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	adminApi := r.Group("/admin")
	{
		adminApi.POST("/vouchers", ctx.Ctl.AdminVoucher.Create)
		adminApi.DELETE("/vouchers/:id", ctx.Ctl.AdminVoucher.Delete)
		adminApi.PUT("/vouchers/:id", ctx.Ctl.AdminVoucher.Update)
		adminApi.GET("/vouchers", ctx.Ctl.AdminVoucher.Get)
<<<<<<< HEAD
		adminApi.GET("/vouchers/user", ctx.Ctl.AdminVoucherUser.Get)
=======
		adminApi.GET("/vouchers/usages", ctx.Ctl.AdminVoucherUsage.Get) //visa
>>>>>>> da0e613 (Initial commit)
	}

	userApi := r.Group("/user")
	{
<<<<<<< HEAD
		userApi.POST("/redemptions", ctx.Ctl.UserRedeemVoucher.Create)
		userApi.GET("/redemptions", nil)

		userApi.POST("/vouchers", nil)
		userApi.GET("/vouchers", nil)
	}

	r.GET("/vouchers", nil)
	r.GET("/vouchers/:id", nil)

=======
		userApi.POST("/redemptions", nil)

		userApi.GET("/redemptions", nil)          //visa
		userApi.GET("/vouchers", nil)             //visa
		userApi.GET("/vouchers/:id", nil)         //visa
		userApi.POST("/vouchers/:id/usages", nil) //visa
		userApi.GET("/vouchers/usages", nil)      //visa
	}

>>>>>>> da0e613 (Initial commit)
	return r
}

// // Endpoints (User)
// GET /user/vouchers
// Fungsi: Mengambil daftar semua voucher yang tersedia untuk pengguna.
// Output: Kumpulan data voucher yang dapat digunakan oleh pengguna.

// GET /user/vouchers/:id
// Fungsi: Mengambil detail informasi dari sebuah voucher tertentu berdasarkan ID-nya.
// Parameter:
// id (path parameter): ID voucher.
// Output: Detail voucher seperti nama, deskripsi, diskon, masa berlaku, dll.
// POST /user/vouchers/:id/usages

// Fungsi: Mencatat penggunaan voucher oleh pengguna (seperti klaim atau transaksi).
// Parameter:
// id (path parameter): ID voucher yang digunakan.
// Input: Informasi terkait transaksi atau klaim voucher.
// Output: Konfirmasi atau status sukses/tidaknya penggunaan voucher.
// GET /user/redemptions

// Fungsi: Mengambil daftar voucher yang telah ditukarkan (redeemed) oleh pengguna.
// Output: Kumpulan data voucher yang sudah digunakan atau diklaim oleh pengguna.
// GET /user/vouchers/usages

// Fungsi: Mengambil log atau riwayat penggunaan voucher oleh pengguna.
// Output: Data riwayat pemakaian voucher, seperti ID voucher, tanggal pemakaian, dan transaksi terkait.
