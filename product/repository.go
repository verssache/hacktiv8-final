package product

import "gorm.io/gorm"

type Repository interface {
	FindAll(product *[]Product) error
	FindByID(id int) (Product, error)
	Store(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(product Product) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(product *[]Product) error {
	err := r.db.Find(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByID(id int) (Product, error) {
	var product Product
	err := r.db.Where("ID = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Store(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Delete(product Product) error {
	err := r.db.Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}
