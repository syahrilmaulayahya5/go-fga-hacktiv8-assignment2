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
