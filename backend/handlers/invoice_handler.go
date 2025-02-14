package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2"

	"github.com/kellymahlangu/invoice-generator/backend/models"
	"github.com/kellymahlangu/invoice-generator/backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(c *gin.Context) {
	invoice := models.Invoice{
		ID: 12345,
		Owner: models.Seller{
			Name:      "John Doe Inc.",
			Address:   "123 Business St, New York, NY",
			Email:     "seller@example.com",
			Contact:   "+1 234 567 890",
			TaxNumber: "123456789",
		},
		Bank: models.BankInfo{
			BankName:          "ABC Bank",
			AccountHolderName: "John Doe Inc.",
			AccountNumber:     12345678,
			BranchCode:        9876,
		},
		Customer: models.Buyer{
			Name:    "Jane Smith",
			Address: "456 Customer Ave, Los Angeles, CA",
			Email:   "customer@example.com",
			Contact: "+1 987 654 321",
		},
		Items: []models.Product{
			{Item: "Product A", Description: "High-quality product", Quantity: 2, UnitPrice: 50.00, Subtotal: 100.00},
			{Item: "Product B", Description: "Another great item", Quantity: 1, UnitPrice: 75.00, Subtotal: 75.00},
		},
		Subtotal:   175.00,
		Tax:        17.50,
		Discounts:  5.00,
		GrandTotal: 187.50,
		DueDate:    time.Now().AddDate(0, 0, 7).Format("January 2, 2006"),
		Status:     "Pending",
		CreatedAt:  time.Now().Format("January 2, 2006"),
	}

	cfg := config.NewBuilder().WithPageSize(pagesize.Type(consts.A4)).Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	// Set Document Header
	m.AddRow(20,
		text.NewCol(12, "INVOICE ( #"+strconv.Itoa(invoice.ID)+" )", props.Text{
			Size:  20,
			Align: align.Type(consts.Center),
			Style: fontstyle.Type(consts.Bold),
		}),
	)
	// Seller & Buyer Information
	// sellerCol := col.New(6)
	// sellerCol.Add(
	// 	text.New("Seller Infomation", props.Text{Size: 12, Style: fontstyle.Type(consts.Bold)}),
	// 	text.New(invoice.Owner.Name),
	// 	text.New(invoice.Owner.Address),
	// 	text.New("Email: "+invoice.Owner.Email),
	// 	text.New("Phone: "+invoice.Owner.Email),
	// 	text.New("Tax Number: "+invoice.Owner.Email),
	// )

	// buyerCol := col.New(6)
	// buyerCol.Add(
	// 	text.New("Buyer Information", props.Text{Size: 12, Style: fontstyle.Type(consts.Bold)}),
	// 	text.New(invoice.Customer.Name),
	// 	text.New(invoice.Customer.Address),
	// 	text.New("Email: "+invoice.Customer.Email),
	// 	text.New("Phone: "+invoice.Customer.Email),
	// )

	sellerAndBuyerRow := utils.GenerateSellerBuyerDetails(invoice.Owner, invoice.Customer)
	m.AddRows(sellerAndBuyerRow...)

	// Invoice ID and Due Date

	dueDateCol := col.New(6)
	dueDateCol.Add(
		text.New("Due Date: " + invoice.DueDate),
	)

	invoiceIdAndDueDateRow := row.New(10).Add(dueDateCol)
	m.AddRows(invoiceIdAndDueDateRow)

	// Add Table to PDF
	rows, err := list.Build(invoice.Items)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Something went wrong converting items!"})
	}

	m.AddRows(rows...)

	// Totals
	totalsRow := row.New(10).Add(
		col.New(12).Add(
			text.New("Subtotal: R"+strconv.FormatFloat(invoice.Subtotal, 'g', 5, 64), props.Text{Align: align.Type(consts.Right)}),
			text.New("Tax: R"+strconv.FormatFloat(invoice.Tax, 'g', 5, 64), props.Text{Align: align.Type(consts.Right)}),
			text.New("Discounts: R"+strconv.FormatFloat(invoice.Discounts, 'g', 5, 64), props.Text{Align: align.Type(consts.Right)}),
			text.New("Grand Total: R"+strconv.FormatFloat(invoice.GrandTotal, 'g', 5, 64), props.Text{
				Size:  14,
				Style: fontstyle.Type(consts.Bold),
				Align: align.Type(consts.Right),
			}),
		),
	)

	m.AddRows(totalsRow)

	// Bank
	bankRow := row.New(10).Add(
		col.New(12).Add(
			text.New("Bank Details", props.Text{Size: 12, Style: fontstyle.Type(consts.Bold)}),
			text.New("Bank Name: "+invoice.Bank.BankName),
			text.New("Account Holder: "+invoice.Bank.AccountHolderName),
			text.New("Account Number: "+strconv.Itoa(invoice.Bank.AccountNumber)),
			text.New("Branch Code: "+strconv.Itoa(invoice.Bank.BranchCode)),
		),
	)

	m.AddRows(bankRow)

	document, err := m.Generate()
	if err != nil {
		log.Fatalf("Failed to generate PDF: %v", err)
	}

	err = document.Save("static/invoices/test.pdf")
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Something went wrong saving invoice!", "err": err.Error()})
	}

	err = document.GetReport().Save("static/reports/test.txt")
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Something went wrong saving invoice report!", "err": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice PDF generated successfully!"})
}

func FetchInvoice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Fetching invoice!"})
}

func DeleteInvoice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Deleting invoice!"})
}

func UpdateInvoice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updating invoice!"})
}

func FetchAllInvoices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Fetching all invoices!"})
}

func DownloadInvoicePdf(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Downloading invoice!"})
}
