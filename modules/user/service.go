package user

import (
	"errors"
	auth "instalasi-pro/middleware/auth"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input UserInput) (User, error)
	Login(input UserInput) (User, error)
	UpdateDataUser(id int, input UpdateUserInput) (User, error)
}

type userService struct {
	repository Repository
}

func NewUserService(repository Repository) *userService {
	return &userService{repository}
}

func (s *userService) RegisterUser(input UserInput) (User, error) {
	user := User{}
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = "customer"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *userService) Login(input UserInput) (string, error) {
	email := input.Email
	password := input.Password
	user, err := s.repository.FindByEmail(email)

	if err != nil || user.ID == 0 {
		return "", errors.New("email or password is incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("email or password is incorrect")
	}

	token, err := auth.NewService().GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetUser(id int) ([]User, error) {
	if id != 0 {
		user, err := s.repository.FindById(id)
		if err != nil {
			return []User{}, err
		}
		return []User{user}, nil
	}
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *userService) UpdateCustomer(id int, customer User) (User, error) {
	customer, err := s.repository.Update(id, customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (s *userService) UpdateDataUser(id int, input UpdateUserInput) (User, error) {

	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Address = input.Address
	user.Phone = input.Phone

	updatedUser, err := s.repository.Update(id, user)

	if err != nil {
		return user, err
	}
	return updatedUser, nil
}
