package invoice

import "time"

type Invoice struct {
	ID            int       `json:"id"`
	OrderID       int       `json:"order_id"`
	Total         float64   `json:"total"`
	Status        string    `json:"status" gorm:"default:'pending'"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CustomerEmail string    `json:"customer_email"`
}
