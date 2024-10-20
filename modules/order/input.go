package order

import "instalasi-pro/modules/user"

type OrderInput struct {
	ProductID int       `json:"product_id" gorm:"product_id"`
	User      user.User `json:"user"`
}

type OrderUpdateInput struct {
	Status       string `json:"status" gorm:"status"`
	TechnicianID int    `json:"technician_id" gorm:"technician_id"`
}
