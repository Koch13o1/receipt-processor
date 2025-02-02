package main

import (
	"receipt-processor/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routes
	r.POST("/receipts/process", controllers.ProcessReceiptHandler)
	r.GET("/receipts/:id/points", controllers.GetReceiptPoints)

	// Start server
	r.Run(":8080")
}
