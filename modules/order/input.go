package order

import "instalasi-pro/modules/user"

type OrderInput struct {
	ProductID int `json:"product_id" form:"product_id"`
	User      user.User
}
