package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/service"
)

type AdminVoucherUsageController struct {
	service service.VoucherService
	logger  *zap.Logger
}

func NewAdminVoucherUsageController(service service.VoucherService, logger *zap.Logger) *AdminVoucherUsageController {
	return &AdminVoucherUsageController{service, logger}
}

func (ctrl *AdminVoucherUsageController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of vouchers"})
}
