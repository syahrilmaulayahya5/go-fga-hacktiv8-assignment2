package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/message"
	"github.com/syahrilmaulayahya5/go-fga-hacktiv8-assignment2/pkg/domain/user"
)

type UserHdlImpl struct {
	userUsecase user.UserUsecase
}

func NewUserHdl(userUseCase user.UserUsecase) user.UserHandler {
	return &UserHdlImpl{userUsecase: userUseCase}
}

func (u *UserHdlImpl) GetUserByEmailHdl(ctx *gin.Context) {
	email := ctx.Query("email")
	result, err := u.userUsecase.GetUserByEmailSvc(ctx, email)
	if err != nil {
		switch err.Error() {
		case "INTERNAL_SERVER_ERROR":
			response := message.NewErrorResponse(http.StatusInternalServerError, "something went wrong")
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return

		case "NOT_FOUND":
			response := message.NewErrorResponse(http.StatusBadRequest, "email not found")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}
	response := message.NewSuccessResponse(result)
	ctx.JSONP(http.StatusOK, response)

}

func (u *UserHdlImpl) InsertUserHdl(ctx *gin.Context) {

	var user user.User

	if err := ctx.ShouldBind(&user); err != nil {
		response := message.NewErrorResponse(http.StatusBadRequest, "failed to bind payload")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if user.Email == "" {
		response := message.NewErrorResponse(http.StatusBadRequest, "email should not be empty")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := u.userUsecase.InsertUserSvc(ctx, user)
	if err != nil {
		response := message.NewErrorResponse(http.StatusBadRequest, "email already registered")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := message.NewSuccessResponse(result)
	ctx.JSONP(http.StatusOK, response)
}
