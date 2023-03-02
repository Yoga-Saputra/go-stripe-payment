package service

import (
	"go-stripe-payment/src/config"
	"go-stripe-payment/src/models"
)

func SavePayment(charge *models.Charge) error {
	if err := config.DB.Create(charge).Error; err != nil {
		return err
	}

	return nil
}
