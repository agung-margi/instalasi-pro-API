package order

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindAll() ([]Order, error)
	FindById(id int) (Order, error)
	FindByUserID(id int) ([]Order, error)
	Create(order Order) (Order, error)
	Update(id int, order Order) (Order, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Order, error) {
	var orders []Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *repository) FindById(id int) (Order, error) {
	var order Order
	err := r.db.Where("id = ?", id).Find(&order).Error
	return order, err
}

func (r *repository) FindByUserID(id int) ([]Order, error) {
	var orders []Order
	err := r.db.Where("user_id = ?", id).Find(&orders).Error
	return orders, err
}

func (r *repository) Create(order Order) (Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repository) Update(id int, order Order) (Order, error) {
	order.ID = id
	err := r.db.Save(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}
