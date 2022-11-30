package category

import "errors"

type Service interface {
	FindAll() ([]Category, error)
	FindByID(id int) (Category, error)
	Store(input CreateCategoryInput) (Category, error)
	Update(ID int, input CreateCategoryInput) (Category, error)
	Delete(ID int) error
	UpdateSoldProduct(ID, Quantity int) (Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Category, error) {
	var categories []Category
	categories, err := s.repository.FindAll()
	if err != nil {
		return categories, err
	}

	if len(categories) == 0 {
		return categories, errors.New("no category found")
	}

	return categories, nil
}

func (s *service) FindByID(id int) (Category, error) {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *service) Store(input CreateCategoryInput) (Category, error) {
	category := Category{}
	category.Type = input.Type

	newCategory, err := s.repository.Store(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) Update(ID int, input CreateCategoryInput) (Category, error) {
	category, err := s.repository.FindByID(ID)
	if err != nil {
		return category, err
	}

	category.Type = input.Type

	updatedCategory, err := s.repository.Update(category)
	if err != nil {
		return updatedCategory, err
	}

	return updatedCategory, nil
}

func (s *service) Delete(ID int) error {
	category, err := s.repository.FindByID(ID)
	if err != nil {
		return err
	}

	err = s.repository.Delete(category)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateSoldProduct(ID, Quantity int) (Category, error) {
	category, err := s.repository.FindByID(ID)
	if err != nil {
		return category, err
	}

	category.SoldProductAmount += Quantity

	updatedCategory, err := s.repository.Update(category)
	if err != nil {
		return updatedCategory, err
	}

	return updatedCategory, nil
}
