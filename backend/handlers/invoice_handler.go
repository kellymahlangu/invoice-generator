package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Creating new invoice!"})
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
