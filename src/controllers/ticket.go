package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ortegasixto7/golang-ticket/src/core/ticket/generate"
)

type TicketController struct {
}

func (ctrl *TicketController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong from tickets",
	})
}

func (ctrl *TicketController) Generate(ctx *gin.Context) {
	request := new(generate.GenerateRequest)

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := generate.Execute(request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ticket created!",
	})
}
