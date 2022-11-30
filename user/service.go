package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetUserByID(ID int) (User, error)
	GetUserByFullName(fullName string) (User, error)
	GetUserByEmail(email string) (User, error)
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	TopUpBalance(ID int, input TopUpBalanceInput) (int, error)
	UpdateBalance(ID int, total int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByFullName(fullName string) (User, error) {
	user, err := s.repository.FindByFullName(fullName)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByEmail(email string) (User, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	var user User
	user.FullName = input.FullName
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = "customer"
	user.Balance = 0

	newUser, err := s.repository.Store(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	var user User
	user.Email = input.Email

	user, err := s.repository.FindByEmail(user.Email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, errors.New("wrong password")
	}

	return user, nil
}

func (s *service) TopUpBalance(ID int, input TopUpBalanceInput) (int, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return 0, err
	}

	if user.ID == 0 {
		return 0, errors.New("user not found")
	}

	if input.Balance < 0 {
		return 0, errors.New("invalid amount")
	}

	total := user.Balance + input.Balance

	if total > 100000000 {
		return 0, errors.New("balance cannot be more than 100.000.000")
	}

	user.Balance = total

	_, err = s.repository.Update(user)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (s *service) UpdateBalance(ID int, total int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	if total < 0 {
		return user, errors.New("invalid amount")
	}

	user.Balance -= total

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
