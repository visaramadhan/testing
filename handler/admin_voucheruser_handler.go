package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/service"
)

type AdminVoucherUserController struct {
	service service.VoucherService
	logger  *zap.Logger
}

func NewAdminVoucherUserController(service service.VoucherService, logger *zap.Logger) *AdminVoucherUserController {
	return &AdminVoucherUserController{service, logger}
}

func (ctrl *AdminVoucherUserController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of vouchers"})
}
