package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/config/gin"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/order"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/server/http/router"
)

type OrderRouterImpl struct {
	ginEngine    engine.HttpServer
	routerGroup  *gin.RouterGroup
	orderHandler order.OrderHandler
}

func NewOrderRouter(ginEngine engine.HttpServer, orderHandler order.OrderHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/v1/order")
	return &OrderRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, orderHandler: orderHandler}
}

func (u *OrderRouterImpl) post() {
	u.routerGroup.POST("/create", u.orderHandler.CreateOrderHdl)

}

func (u *OrderRouterImpl) Routers() {
	u.post()
}
