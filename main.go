package main

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	engine "github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/config/gin"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/config/postgres"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/message"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/user"
	v1 "github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/server/http/router/v1"
	"gorm.io/gorm"

	userrepo "github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/repository/user"
	userhandler "github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/server/http/handler/user"

	userusecase "github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/usecase/user"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
func main() {
	postgresCln := postgres.NewPostgresConnecion(postgres.Config{
		Host:         "localhost",
		Port:         "5432",
		User:         "postgres",
		Password:     "password",
		DatabaseName: "postgres",
	})
	dbMigrate(postgresCln.GetClient())
	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})
	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger())

	startTime := time.Now()
	ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
		// secara default map jika di return dalam
		// response API, dia akan menjadi JSON
		response := message.BaseResponse{}
		response.Meta.Status = http.StatusOK
		response.Meta.Message = "server up and running"
		response.Meta.StartTime = &startTime
		response.Data = nil

		ctx.JSON(http.StatusOK, response)
	})
	userRepo := userrepo.NewUserRepo(postgresCln)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	userHandler := userhandler.NewUserHdl(userUsecase)
	v1.NewUserRouter(ginEngine, userHandler).Routers()

	ginEngine.Serve()
}
