package product

import "errors"

type Service interface {
	FindAll() ([]Product, error)
	FindByID(id int) (Product, error)
	Store(input CreateProductInput) (Product, error)
	Update(ID int, input CreateProductInput) (Product, error)
	Delete(ID int) error
	UpdateStock(ID, Quantity int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Product, error) {
	var products []Product
	err := s.repository.FindAll(&products)
	if err != nil {
		return products, err
	}

	if len(products) == 0 {
		return products, errors.New("no product found")
	}

	return products, nil
}

func (s *service) FindByID(id int) (Product, error) {
	product, err := s.repository.FindByID(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) Store(input CreateProductInput) (Product, error) {
	product := Product{}
	product.Title = input.Title
	product.Price = input.Price
	product.Stock = input.Stock
	product.CategoryID = input.CategoryID

	if product.CategoryID == 0 {
		return product, errors.New("category id is required")
	}

	if product.Price == 0 || product.Price < 0 {
		return product, errors.New("price is required")
	}

	if product.Price > 50000000 {
		return product, errors.New("price is not valid")
	}

	if product.Stock < 5 {
		return product, errors.New("stock must be greater than 5")
	}

	newProduct, err := s.repository.Store(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) Update(ID int, input CreateProductInput) (Product, error) {
	product, err := s.repository.FindByID(ID)
	if err != nil {
		return product, err
	}

	product.Title = input.Title
	product.Price = input.Price
	product.Stock = input.Stock
	product.CategoryID = input.CategoryID

	if product.CategoryID == 0 {
		return product, errors.New("category id is required")
	}

	if product.Price == 0 || product.Price < 0 {
		return product, errors.New("price is required")
	} else if product.Price > 50000000 {
		return product, errors.New("price is not valid")
	}

	if product.Stock < 5 {
		return product, errors.New("stock must be greater than 5")
	}

	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}

func (s *service) Delete(ID int) error {
	product, err := s.repository.FindByID(ID)
	if err != nil {
		return err
	}

	if product.ID == 0 {
		return errors.New("product not found")
	}

	err = s.repository.Delete(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateStock(ID, Quantity int) (Product, error) {
	product, err := s.repository.FindByID(ID)
	if err != nil {
		return product, err
	}

	product.Stock -= Quantity

	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}
