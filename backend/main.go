package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kellymahlangu/invoice-generator/backend/routes"
)

func main() {
	router := gin.Default()

	routes.SetupInvoiceRoutes(router)
	router.Run(":8080")
}
