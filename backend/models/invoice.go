package models

import "time"

type Invoice struct {
	ID         int
	Owner      Seller
	Bank       BankInfo
	Customer   Buyer
	Items      []Product
	Subtotal   float64
	Tax        float64
	Discounts  float64
	GrandTotal float64
	DueDate    time.Time
	Status     string
	CreatedAt  time.Time
	Notes      []string
}
