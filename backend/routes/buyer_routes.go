package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kellymahlangu/invoice-generator/backend/handlers"
)

func SetupBuyerRoutes(router *gin.Engine) {
	buyerGroup := router.Group("/buyers")
	{
		buyerGroup.GET("/", handlers.FetchAllBuyers)
		buyerGroup.POST("/", handlers.NewBuyer)
		buyerGroup.GET("/:id", handlers.FetchBuyer)
		buyerGroup.GET("/:id/invoices", handlers.FetchInvoicesForBuyer)
		buyerGroup.DELETE("/:id", handlers.DeleteBuyer)
		buyerGroup.PUT("/:id", handlers.UpdateBuyer)
	}
}
