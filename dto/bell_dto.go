package dto

import "time"

type BellDTO struct {
	ID           int       `json:"id"`
	StoreID      int       `json:"store_id"`
	DeviceID     string    `json:"device_id"`
	Status       string    `json:"status"`
	LastCalledAt time.Time `json:"last_called_at"`
	StoreName    string    `json:"store_name"`
}
