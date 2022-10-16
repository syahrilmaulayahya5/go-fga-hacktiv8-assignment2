package order

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/message"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/order"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/request"
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

func (u *OrderHdlImpl) GetOrderByUserIdHdl(ctx *gin.Context) {

	id, inputErr := strconv.ParseUint(ctx.Query("user_id"), 10, 64)
	if inputErr != nil {
		response := message.NewErrorResponse(http.StatusBadRequest, "user_id must be int with value > 0")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := u.orderUsecase.GetOrderByUserIdSvc(ctx, uint(id))
	if err != nil {
		response := message.NewErrorResponse(http.StatusBadRequest, "something went wrong")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := message.NewSuccessResponse(result)
	ctx.JSONP(http.StatusOK, response)
}

func (u *OrderHdlImpl) UpdateOrderByIdHdl(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		response := message.NewErrorResponse(http.StatusBadRequest, "id must int with value > 0")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var input request.UpdateOrderRequest

	if err := ctx.ShouldBind(&input); err != nil {
		response := message.NewErrorResponse(http.StatusBadRequest, "failed to bind payload")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := u.orderUsecase.UpdateOrderByIdSvc(ctx, uint(id), request.ToDomain(input))
	if err != nil {
		response := message.NewErrorResponse(http.StatusBadRequest, "order id not found")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := message.NewSuccessResponse(result)
	ctx.JSONP(http.StatusOK, response)
}

func (u *OrderHdlImpl) DeleteOrderByIdHdl(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		response := message.NewErrorResponse(http.StatusBadRequest, "id must int with value > 0")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result := u.orderUsecase.DeleteOrderByIdSvc(ctx, uint(id))
	if err != nil {
		response := message.NewErrorResponse(http.StatusBadRequest, "order id not found")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := message.NewSuccessResponse(result)
	ctx.JSONP(http.StatusOK, response)
}
