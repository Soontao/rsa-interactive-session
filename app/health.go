package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *WebApplication) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": http.StatusOK, "Service": app.param.ServiceName})
}
