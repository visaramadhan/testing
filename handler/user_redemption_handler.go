package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/service"
)

type UserRedeemVoucherController struct {
	service service.RedeemVoucherService
	logger  *zap.Logger
}

func NewUserRedeemVoucherController(service service.RedeemVoucherService, logger *zap.Logger) *UserRedeemVoucherController {
	return &UserRedeemVoucherController{service, logger}
}

func (ctrl *UserRedeemVoucherController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Redemption created"})
}
