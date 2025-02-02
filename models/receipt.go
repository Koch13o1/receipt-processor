package models

import (
	"errors"
	"strconv"
)

// Item represents an item in a receipt.
// I have added a PriceValue field that convertes the Price from string to PriceValue, i.e. float64.
type Item struct {
	ShortDescription string  `json:"shortDescription" binding:"required"`
	Price            string  `json:"price" binding:"required"`
	PriceValue       float64 `json:"-"`
}

// Receipt represents a receipt object.
// I have added a TotalValue field wherein the Total string is parsed and converted to TotalValue which is float64.
type Receipt struct {
	Retailer     string  `json:"retailer" binding:"required"`
	PurchaseDate string  `json:"purchaseDate" binding:"required"`
	PurchaseTime string  `json:"purchaseTime" binding:"required"`
	Items        []Item  `json:"items" binding:"required,dive"`
	Total        string  `json:"total" binding:"required"`
	TotalValue   float64 `json:"-"`
}

// ValidateAndConvert function is used to convert string total and price fields of items to float64
// It also ensures if sum of prices is equal to total, to validate the receipt.

func (r *Receipt) ValidateAndConvert() error {
	total, err := strconv.ParseFloat(r.Total, 64)

	var totalSum float64

	if err != nil {
		return err
	}
	r.TotalValue = total

	for i := range r.Items {
		price, err := strconv.ParseFloat(r.Items[i].Price, 64)
		if err != nil {
			return err
		}
		r.Items[i].PriceValue = price
		totalSum += price
	}

	if totalSum != r.TotalValue {
		return errors.New("receipt total does not match sum of item prices")
	}

	return nil
}
