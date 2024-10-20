package technician

import "instalasi-pro/modules/user"

type Service interface {
	FindAll() ([]user.User, error)
	FindById(id int) (user.User, error)
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
