package models

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
	DueDate    string
	Status     string
	CreatedAt  string
	Notes      []string
}
