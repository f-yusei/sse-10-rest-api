package dto

// StoreDTO はStoreエンティティのデータ転送オブジェクトです。
type StoreDTO struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	DisplayMessage string `json:"display_message"`
}
