package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kellymahlangu/invoice-generator/backend/handlers"
)

func SetupInvoiceRoutes(router *gin.Engine) {
	invoiceGroup := router.Group("/invoices")
	{
		invoiceGroup.GET("/", handlers.FetchAllInvoices)
		invoiceGroup.POST("/", handlers.CreateInvoice)
		invoiceGroup.GET("/:id", handlers.FetchInvoice)
		invoiceGroup.GET("/:id/pdf", handlers.DownloadInvoicePdf)
		invoiceGroup.DELETE("/:id", handlers.DeleteInvoice)
		invoiceGroup.PUT("/:id", handlers.FetchInvoice)
	}
}
