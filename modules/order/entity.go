package order

import (
	"instalasi-pro/modules/product"
	"instalasi-pro/modules/user"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           int `json:"id" gorm:"primaryKey"`
	UserID       int `json:"user_id" gorm:"not null"`
	TechnicianID int `json:"technician_id" gorm:"not null"`
	ProductID    int `json:"product_id" gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt

	Product    product.Product `gorm:"foreignKey:ProductID"`
	User       user.User       `gorm:"foreignKey:UserID"`
	Technician user.User       `gorm:"foreignKey:TechnicianID"`
}
