package model

import (
	"time"
)

type Bell struct {
	ID           int
	StoreID      int
	DeviceID     string
	Status       string
	LastCalledAt time.Time
}
