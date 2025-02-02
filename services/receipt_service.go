package services

import (
	"errors"
	"math"
	"receipt-processor/models"
	"receipt-processor/storage"
	"receipt-processor/utils"
	"strings"
	"time"
)

// SaveReceipt stores the receipt and returns a unique ID
func SaveReceipt(receipt models.Receipt) (string, error) {
	// Validate and convert prices
	if err := receipt.ValidateAndConvert(); err != nil {
		return "", errors.New("invalid receipt: " + err.Error())
	}

	// Generate unique ID and store receipt
	id := utils.GenerateUUID()
	storage.StoreReceipt(id, receipt)

	return id, nil
}

// GetReceiptPoints retrieves the points assigned to a receipt
func GetReceiptPoints(id string) (int, bool) {
	receipt, exists := storage.GetReceipt(id)
	if !exists {
		return 0, false
	}

	points := 0

	// Retailer name points: One point for every alphanumeric character
	points += countAlphanumeric(receipt.Retailer)

	// 50 points if total is a round number
	if isRoundNumber(receipt.TotalValue) {
		points += 50
	}

	// 25 points if total is a multiple of 0.25
	if isMultipleOfQuarter(receipt.TotalValue) {
		points += 25
	}

	// 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Item description length + price rule
	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			pricePoints := int(math.Ceil(item.PriceValue * 0.2))
			points += pricePoints
		}
	}

	// 6 points if purchase date is an odd day
	if isOddDay(receipt.PurchaseDate) {
		points += 6
	}

	// 10 points if purchase time is between 14:00 - 16:00
	if isBetweenTwoAndFour(receipt.PurchaseTime) {
		points += 10
	}

	return points, true
}

// countAlphanumeric function counts the number of alphanumeric characters in a string
func countAlphanumeric(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			count++
		}
	}
	return count
}

// isRoundNumber function checks if the total is a whole number with no cents
func isRoundNumber(total float64) bool {
	return math.Mod(total, 1.0) == 0
}

// isMultipleOfQuarter function checks if the total is a multiple of 0.25
func isMultipleOfQuarter(total float64) bool {
	return math.Mod(total, 0.25) == 0
}

// isOddDay function checks if the day in the purchase date is odd
func isOddDay(dateStr string) bool {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return false
	}
	return date.Day()%2 != 0
}

// isBetweenTwoAndFour function checks if the purchase time is between 14:00 and 16:00
func isBetweenTwoAndFour(timeStr string) bool {
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return false
	}
	return parsedTime.Hour() >= 14 && parsedTime.Hour() < 16
}
