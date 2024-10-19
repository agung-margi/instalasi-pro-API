package user

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindAll() ([]User, error)
	FindById(id int) (User, error)
	FindByEmail(email string) (User, error)
	Save(user User) (User, error)
	Update(id int, user User) (User, error)
	Delete(id int) error
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindById(id int) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).Find(&user).Error
	return user, err
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(id int, user User) (User, error) {
	user.ID = id
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Delete(id int) error {
	err := r.db.Where("id = ?", id).Delete(&User{}).Error
	return err
}
