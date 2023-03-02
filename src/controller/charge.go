package controller

import (
	"go-stripe-payment/src/models"
	"go-stripe-payment/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func Payment(c *gin.Context) {
	var payment models.Charge
	c.BindJSON(&payment)

	apiKey := "sk_test_51Mh24KLyXPxRnDeXv2oxcH8W2kMOmrOSBVtmgiGpAhL458YNddPPYmzK1euMaozF3Dx7ItaDO9G71QMxgFTlo79c00xWH3CCqj"
	stripe.Key = apiKey

	data, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(payment.Amount),
		Currency:     stripe.String(string(stripe.CurrencyIDR)),
		Description:  stripe.String(payment.ProductName),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String(payment.ReceiptEmail),
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err1 := service.SavePayment(&payment)
	if err1 != nil {
		c.JSON(http.StatusNotFound, "error occured")
		return
	}

	c.JSON(http.StatusOK, data)
}
