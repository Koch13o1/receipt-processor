package models

import (
	"errors"
	"strconv"
)

// Item represents an item in a receipt.
// Price remains a string (as received from JSON), and PriceValue holds the float64 value.
type Item struct {
	ShortDescription string  `json:"shortDescription" binding:"required"`
	Price            string  `json:"price" binding:"required"`
	PriceValue       float64 `json:"-"`
}

// Receipt represents a receipt.
// Total is received as a string, and TotalValue holds the float64 value.
// Points is computed when the receipt is saved.
type Receipt struct {
	Retailer     string  `json:"retailer" binding:"required"`
	PurchaseDate string  `json:"purchaseDate" binding:"required"`
	PurchaseTime string  `json:"purchaseTime" binding:"required"`
	Items        []Item  `json:"items" binding:"required,dive"`
	Total        string  `json:"total" binding:"required"`
	TotalValue   float64 `json:"-"`
	Points       int     `json:"-"`
}

// ValidateAndConvert converts the Total and each item's Price from string to float64.
// This version does not check if the sum of item prices equals the total.
func (r *Receipt) ConvertToFloat64() error {
	total, err := strconv.ParseFloat(r.Total, 64)
	if err != nil {
		return errors.New("invalid total value")
	}
	r.TotalValue = total

	for i := range r.Items {
		price, err := strconv.ParseFloat(r.Items[i].Price, 64)
		if err != nil {
			return errors.New("invalid item price")
		}
		r.Items[i].PriceValue = price
	}

	return nil
}
