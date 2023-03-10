package models

import "gorm.io/gorm"

type Charge struct {
	gorm.Model
	Amount       int64  `json:"amount"`
	ReceiptEmail string `json:"receiptMail"`
	ProductName  string `json:"productName"`
}

func (c *Charge) TableName() string {
	return "charge"
}
