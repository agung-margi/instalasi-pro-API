package invoice

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(invoice *Invoice) error
	UpdateStatus(id int, status string) (Invoice, error)
	FindById(id int) (Invoice, error)
	FindAll() ([]Invoice, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(invoice Invoice) (Invoice, error) {
	err := r.db.Create(&invoice).Error
	if err != nil {
		return invoice, err
	}
	return invoice, nil
}

func (r *repository) UpdateStatus(id int, status string) (Invoice, error) {
	var invoice Invoice
	err := r.db.Model(&invoice).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return invoice, err
	}
	return invoice, nil
}

func (r *repository) FindById(id int) (Invoice, error) {
	var invoice Invoice
	err := r.db.Where("id = ?", id).Find(&invoice).Error
	if err != nil {
		return invoice, err
	}
	return invoice, nil
}

func (r *repository) FindAll() ([]Invoice, error) {
	var invoices []Invoice
	err := r.db.Find(&invoices).Error
	if err != nil {
		return invoices, err
	}
	return invoices, nil
}
