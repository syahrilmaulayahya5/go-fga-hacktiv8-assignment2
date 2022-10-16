package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/config/gin"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/user"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/server/http/router"
)

type UserRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	userHandler user.UserHandler
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/v1/user")
	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler}
}

func (u *UserRouterImpl) get() {
	u.routerGroup.GET("/get", u.userHandler.GetUserByEmailHdl)

}

func (u *UserRouterImpl) Routers() {
	u.get()
}
