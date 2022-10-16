package order

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	CreateOrderHdl(ctx *gin.Context)
	GetOrderByUserIdHdl(ctx *gin.Context)
	UpdateOrderByIdHdl(ctx *gin.Context)
	DeleteOrderByIdHdl(ctx *gin.Context)
}
