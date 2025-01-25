package entity

import "time"

type CallLog struct {
	ID       int       `json:"id" gorm:"primaryKey;autoIncrement"`
	BellID   int       `json:"bell_id" gorm:"not null;index"`                                  // 呼び出されたベルID
	StoreID  int       `json:"store_id" gorm:"not null;index"`                                 // 呼び出した店舗ID
	CalledAt time.Time `json:"called_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`      // 呼び出しの発生時刻
	Status   string    `json:"status" gorm:"type:enum('active','completed');default:'active'"` // 呼び出しの状態

	// リレーション
	Bell  Bell  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:BellID;references:ID"`
	Store Store `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:StoreID;references:ID"`
}
