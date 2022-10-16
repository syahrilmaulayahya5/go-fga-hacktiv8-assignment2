package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/message"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/order"
)

type OrderHdlImpl struct {
	orderUsecase order.OrderUsecase
}

func NewOrderHandler(orderUsecase order.OrderUsecase) order.OrderHandler {
	return &OrderHdlImpl{orderUsecase: orderUsecase}
}

func (u *OrderHdlImpl) CreateOrderHdl(ctx *gin.Context) {

	var input order.Order

	if err := ctx.ShouldBind(&input); err != nil {
		response := message.NewErrorResponse(http.StatusBadRequest, "failed to bind payload")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := u.orderUsecase.CreateOrderSvc(ctx, input)
	if err != nil {
		response := message.NewErrorResponse(http.StatusBadRequest, "user id not found")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := message.NewSuccessResponse(result)
	ctx.JSONP(http.StatusOK, response)
}
