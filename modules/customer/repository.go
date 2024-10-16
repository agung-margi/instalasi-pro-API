package customer

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetAll() ([]Customer, error)
	GetById(id int) (Customer, error)
	GetByEmail(email string) (Customer, error)
	Save(customer Customer) (Customer, error)
	Update(id string, customer Customer) (Customer, error)
	Delete(id string) error
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Customer, error) {
	var customers []Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *repository) GetById(id int) (Customer, error) {
	var customer Customer
	err := r.db.Where("id = ?", id).Find(&customer).Error
	return customer, err
}

func (r *repository) Save(customer Customer) (Customer, error) {
	err := r.db.Create(&customer).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *repository) Update(id string, customer Customer) (Customer, error) {
	err := r.db.Save(&customer).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *repository) Delete(id string) error {
	err := r.db.Where("id = ?", id).Delete(&Customer{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByEmail(email string) (Customer, error) {
	var customer Customer
	err := r.db.Where("email = ?", email).Find(&customer).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}
