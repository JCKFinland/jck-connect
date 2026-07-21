package wallet

import (
	"github.com/gin-gonic/gin"

	wallethandler "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/handler"
)

func RegisterRoutes(
	api *gin.RouterGroup,
	handler *wallethandler.Handler,
) {

	wallet := api.Group("/wallet")
	{
		wallet.GET("", handler.GetWallet)
	}
}
