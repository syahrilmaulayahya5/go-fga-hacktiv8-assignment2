package order

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	CreateOrderHdl(ctx *gin.Context)
}
