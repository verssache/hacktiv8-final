package category

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Category, error)
	FindByID(id int) (Category, error)
	Store(category Category) (Category, error)
	Update(category Category) (Category, error)
	Delete(category Category) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Preload("Products").Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (r *repository) FindByID(id int) (Category, error) {
	var category Category
	err := r.db.Preload("Products").Where("ID = ?", id).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repository) Store(category Category) (Category, error) {
	err := r.db.Preload("Products").Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Update(category Category) (Category, error) {
	err := r.db.Preload("Products").Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Delete(category Category) error {
	err := r.db.Preload("Products").Delete(&category).Error
	if err != nil {
		return err
	}

	return nil
}
