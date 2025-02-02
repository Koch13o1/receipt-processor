package controllers

import (
	"net/http"
	"receipt-processor/models"
	"receipt-processor/services"

	"github.com/gin-gonic/gin"
)

// ProcessReceipt function is used to handle the submitted receipt through the POST request
// It assigns a unique ID to every receipt submission
func ProcessReceiptHandler(c *gin.Context) {
	var receipt models.Receipt

	// Bind JSON
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt format. Please verify input."})
		return
	}

	// Save receipt with validation
	id, err := services.SaveReceipt(receipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetReceiptPoints function checks if a receipt with the ID exists.
// If it does exists it returns the calculated points.
func GetReceiptPoints(c *gin.Context) {
	id := c.Param("id")

	points, exists := services.GetReceiptPoints(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt ID not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
