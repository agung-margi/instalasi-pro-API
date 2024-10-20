package order

import (
	"instalasi-pro/modules/product"
	"instalasi-pro/modules/user"
	"time"
)

type Order struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	UserID       int    `json:"user_id" gorm:"not null"`
	TechnicianID int    `json:"technician_id" gorm:"null"`
	ProductID    int    `json:"product_id" gorm:"not null"`
	Status       string `json:"status" gorm:"default:'pending'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Product    product.Product `gorm:"foreignKey:ProductID"`
	User       user.User       `gorm:"foreignKey:UserID"`
	Technician user.User       `gorm:"foreignKey:TechnicianID"`
}
