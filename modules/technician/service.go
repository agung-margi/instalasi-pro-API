package technician

import (
	"instalasi-pro/modules/user"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	FindAll() ([]user.User, error)
	FindById(id int) (user.User, error)
	Register(user TechnicianInput) (user.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]user.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *service) FindById(id int) (user.User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) RegisterTechnician(input TechnicianInput) (user.User, error) {
	user := user.User{}
	user.Email = input.Email
	user.Name = input.Name
	user.Address = input.Address
	user.Phone = input.Phone

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = "technician"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
