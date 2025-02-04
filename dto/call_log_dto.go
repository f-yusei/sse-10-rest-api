package dto

import "time"

type CallLogDTO struct {
	ID       int       `json:"id"`
	BellID   int       `json:"bell_id"`
	StoreID  int       `json:"store_id"`
	CalledAt time.Time `json:"called_at"`
	Status   string    `json:"status"`
}
