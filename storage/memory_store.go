package storage

import "receipt-processor/models"

var receiptStore = make(map[string]models.Receipt)

// StoreReceipt saves a receipt under a given ID.
func StoreReceipt(id string, receipt models.Receipt) {
	receiptStore[id] = receipt
}

// GetReceipt fetches a receipt by ID.
func GetReceipt(id string) (models.Receipt, bool) {
	receipt, exists := receiptStore[id]
	return receipt, exists
}
