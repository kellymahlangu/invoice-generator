package utils

import (
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/kellymahlangu/invoice-generator/backend/models"
)

func UniqueInvoiceNumber() {

}

func GenerateSellerBuyerDetails(seller models.Seller, buyer models.Buyer) []core.Row {
	// sellerHeader := text.NewCol(6, "Seller Infomation", props.Text{Size: 12, Style: consts.Bold})
	// buyerHeader := text.NewCol(6, "Buyer Infomation", props.Text{Size: 12, Style: consts.Bold})
	// headerRow := row.New(10).Add(sellerHeader, buyerHeader)

	sellerName := text.NewCol(6, seller.Name)
	buyerName := text.NewCol(6, buyer.Name)
	nameRow := row.New(5).Add(sellerName, buyerName)

	// Address
	sellerAdd := text.NewCol(6, seller.Address)
	buyerAdd := text.NewCol(6, buyer.Address)
	addRow := row.New(5).Add(sellerAdd, buyerAdd)
	// Email
	sellerEmail := text.NewCol(6, "Email :"+seller.Email)
	buyerEmail := text.NewCol(6, "Email :"+buyer.Email)
	emailRow := row.New(5).Add(sellerEmail, buyerEmail)
	// Phone
	sellerCont := text.NewCol(6, "Phone :"+seller.Contact)
	buyerCont := text.NewCol(6, "Phone :"+buyer.Contact)
	contRow := row.New(5).Add(sellerCont, buyerCont)
	// Tax Number
	sellerTax := text.NewCol(6, "Tax Number :"+seller.TaxNumber)
	buyerTax := text.NewCol(6, "")
	taxRow := row.New(5).Add(sellerTax, buyerTax)

	return []core.Row{
		// headerRow,
		nameRow,
		addRow,
		emailRow,
		contRow,
		taxRow,
	}
}
