package main

import (
	"fmt"
	"go-stripe-payment/src/config"
	"go-stripe-payment/src/models"
	"go-stripe-payment/src/router"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err)
	}
	config.DB.AutoMigrate(&models.Charge{})
	r := router.ChargeRouter()
	r.Run()
}
