package signup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ortegasixto7/golang-ticket/src/integration/repository"
)

type Controller struct {
	Repo repository.IntegrationRepositoryInterface
}

func (ctrl *Controller) SignUp(ctx *gin.Context) {
	request := new(Request)

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := Handle(request, ctrl.Repo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
