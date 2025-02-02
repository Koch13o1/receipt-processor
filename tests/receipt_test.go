package tests

// To test if we receive a HTTP status 200 OK when we POST a receipt

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipt-processor/controllers"
	"receipt-processor/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProcessReceipt(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/receipts/process", controllers.ProcessReceiptHandler)

	receipt := models.Receipt{
		Retailer:     "M&M Market",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "6.49",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		},
	}

	body, _ := json.Marshal(receipt)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
