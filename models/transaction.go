package models

type Transaction struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	UserID uint    `json:"user_id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"` // "withdraw" or "deposit"
}
