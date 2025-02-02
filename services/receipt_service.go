package services

import (
	"math"
	"regexp"
	"strings"
	"time"

	"receipt-processor/models"
	"receipt-processor/storage"
	"receipt-processor/utils"
)

// SaveReceipt computes the points, stores it, and returns a unique ID.
func SaveReceipt(receipt models.Receipt) (string, error) {
	// Validate and convert string values to float64
	if err := receipt.ConvertToFloat64(); err != nil {
		return "", err
	}

	// Calculate points based on the rules and assign to the receipt
	receipt.Points = CalculatePoints(receipt)

	// Generate a unique ID and store the receipt (including its points)
	id := utils.GenerateUUID()
	storage.StoreReceipt(id, receipt)

	return id, nil
}

// GetReceiptPoints retrieves the precomputed points for a receipt.
func GetReceiptPoints(id string) (int, bool) {
	receipt, exists := storage.GetReceipt(id)
	if !exists {
		return 0, false
	}
	return receipt.Points, true
}

// CalculatePoints applies all the rules to compute the points for the receipt.
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// One point for every alphanumeric character in the retailer name.
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(re.FindAllString(receipt.Retailer, -1))

	// 50 points if the total is a round dollar amount (no cents).
	if math.Mod(receipt.TotalValue, 1.0) == 0 {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if math.Mod(receipt.TotalValue, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items.
	points += (len(receipt.Items) / 2) * 5

	// For each item, if the trimmed description length is a multiple of 3,
	// add ceil(price * 0.2) points.
	for _, item := range receipt.Items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		if len(trimmed)%3 == 0 {
			// Calculate points for this item.
			itemPoints := int(math.Ceil(item.PriceValue * 0.2))
			points += itemPoints
		}
	}

	// 6 points if the purchase date's day is odd.
	if t, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil {
		if t.Day()%2 == 1 {
			points += 6
		}
	}

	// 10 points if the purchase time is after 2:00 PM and before 4:00 PM.
	if t, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
		hour := t.Hour()
		if hour >= 14 && hour < 16 {
			points += 10
		}
	}

	// if receipt.TotalValue > 10.00 {
	// 	points += 5
	// }

	return points
}
