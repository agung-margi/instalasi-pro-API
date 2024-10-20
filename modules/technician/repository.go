package technician

import (
	"instalasi-pro/modules/user"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindAll() ([]user.User, error)
	FindById(id int) (user.User, error)
	Save(user user.User) (user.User, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]user.User, error) {
	var users []user.User
	err := r.db.Where("role = ?", "technician").Find(&users).Error
	return users, err
}

func (r *repository) FindById(id int) (user.User, error) {
	var user user.User
	err := r.db.Where("role = ?", "technician").Where("id = ?", id).Find(&user).Error
	return user, err
}

func (r *repository) Save(user user.User) (user.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
