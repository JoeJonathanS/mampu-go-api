package models

type User struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name"`
	Email   string  `json:"email" gorm:"unique"`
	Balance float64 `json:"balance"`
}
