package models

import "time"

type OrderHistory struct {
	ID         string     `json:"id"`
	Items      []CartItem `json:"items"`
	TotalItem  int        `json:"total_item"`
	TotalPrice int        `json:"total_price"`
	Payment    int        `json:"payment"`
	Change     int        `json:"change"`
	CreatedAt  time.Time  `json:"created_at"`
}
