package services

import (
	"encoding/json"
	"os"
	"time"
	"weekly3/models"

	"github.com/google/uuid"
)

func OrderHistory(items []models.CartItem, totalItem, totalPrice, payment, change int) error {

	const filePath = "data/history.json"

	var histories []models.OrderHistory

	// Buat folder data jika belum ada
	if err := os.MkdirAll("data", 0755); err != nil {
		return err
	}

	// Jika file belum ada, buat dengan array kosong
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte("[]"), 0644); err != nil {
			return err
		}
	}

	// Baca isi file
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Parse JSON
	if len(file) > 0 {
		if err := json.Unmarshal(file, &histories); err != nil {
			return err
		}
	}

	// Tambahkan histori baru
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

	// Simpan kembali ke file
	data, err := json.MarshalIndent(histories, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
