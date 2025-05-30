package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ortegasixto7/golang-ticket/src/user/repository"
)

type Controller struct {
	repo repository.UserRepositoryInterface
}

func NewController(repo repository.UserRepositoryInterface) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (ctrl *Controller) Handle(ctx *gin.Context) {
	request := new(Request)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := Execute(ctx.Request.Context(), request, ctrl.repo)
	if err != nil {
		switch err.(type) {
		case *RequestValidationError:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			if err == ErrInvalidCredentials {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}
