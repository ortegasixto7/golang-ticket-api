package regenerate_token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ortegasixto7/golang-ticket/src/app/repository"
)

type Controller struct {
	Repo repository.AppRepositoryInterface
}

func (ctrl *Controller) Handle(ctx *gin.Context) {
	request := new(Request)

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := Execute(ctx, request, ctrl.Repo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
