package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewBuyer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Creating new Buyer!"})
}

func UpdateBuyer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updating Buyer!"})
}

func DeleteBuyer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Deleting Buyer!"})
}

func FetchBuyer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Fetching Buyer!"})
}

func FetchAllBuyers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Fetching all Buyers!"})
}

func FetchInvoicesForBuyer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Fetching all Buyers!"})
}
