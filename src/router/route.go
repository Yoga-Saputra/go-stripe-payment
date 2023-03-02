package router

import (
	"go-stripe-payment/src/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChargeRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
	group := r.Group("/stripe")
	{
		group.POST("payment", controller.Payment)
	}
	return r
}
