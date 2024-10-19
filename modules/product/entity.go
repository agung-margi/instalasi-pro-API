package product

import "time"

type Product struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"varchar(100) not null"`
	Description string `json:"description" gorm:"varchar(255) not null"`
	Price       int    `json:"price" gorm:"int not null"`
	IsActive    bool   `json:"status" gorm:"default:true"`
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
}
