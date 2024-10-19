package product

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindAll() ([]Product, error)
	FindById(id int) (Product, error)
	Save(product Product) (Product, error)
	Update(id int, product Product) (Product, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *repository) FindById(id int) (Product, error) {
	var product Product
	err := r.db.Where("id = ?", id).Find(&product).Error
	return product, err
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *repository) Update(id int, product Product) (Product, error) {
	product.ID = id
	err := r.db.Save(&product).Error
	return product, err
}
