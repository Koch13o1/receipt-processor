package controllers

import (
	"net/http"
	"receipt-processor/models"
	"receipt-processor/services"

	"github.com/gin-gonic/gin"
)

// ProcessReceiptHandler handles receipt submission via POST,
// validating the receipt, computing its points, and then saving it.
func ProcessReceiptHandler(c *gin.Context) {
	var receipt models.Receipt

	// Bind JSON into the receipt struct
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format. Please verify input."})
		return
	}

	// Save the receipt (which will validate, compute points, and store it)
	id, err := services.SaveReceipt(receipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond with the generated ID
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetReceiptPoints handles GET requests and returns the precomputed points.
func GetReceiptPoints(c *gin.Context) {
	id := c.Param("id")

	points, exists := services.GetReceiptPoints(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt ID not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
