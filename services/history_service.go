package services

import (
	"encoding/json"
	"os"
	"time"
	"weekly3/models"

	"github.com/google/uuid"
)

func OrderHistory(items []models.CartItem, totalItem, totalPrice, payment, change int) error {
	var histories []models.OrderHistory

	file, err := os.ReadFile("data/history.json")
	if err == nil && len(file) > 0 {
		json.Unmarshal(file, &histories)
	}

	history := models.OrderHistory{
		ID:         uuid.NewString(),
		Items:      items,
		TotalItem:  totalItem,
		TotalPrice: totalPrice,
		Payment:    payment,
		Change:     change,
		CreatedAt:  time.Now(),
	}

	histories = append(histories, history)

	data, err := json.MarshalIndent(histories, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile("data/history.json", data, 0644)
}
