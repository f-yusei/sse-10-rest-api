package entity

import "time"

type Bell struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreID      int       `json:"store_id" gorm:"not null;index"`
	DeviceID     string    `json:"device_id" gorm:"type:varchar(255);not null"`
	Status       string    `json:"status" gorm:"type:enum('idle','calling');default:'idle'"`
	LastCalledAt time.Time `json:"last_called_at" gorm:"type:timestamp"`
	Store        Store     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:StoreID;references:ID"`
}
