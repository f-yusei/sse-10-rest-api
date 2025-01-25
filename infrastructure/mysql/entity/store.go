package entity

type Store struct {
	ID             int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string `json:"name" gorm:"type:varchar(255);not null"`
	Bells          []Bell `gorm:"foreignKey:StoreID"`
	DisplayMessage string `json:"display_message" gorm:"type:varchar(255);not null;default:'呼び出し中'"` // 呼び出し時に表示するメッセージ
}
