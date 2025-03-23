package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
}

func (ctrl *TicketController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong from tickets",
	})
}
