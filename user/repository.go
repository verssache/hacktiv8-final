package user

import "gorm.io/gorm"

type Repository interface {
	FindByID(id int) (User, error)
	FindByFullName(fullName string) (User, error)
	FindByEmail(email string) (User, error)
	Store(user User) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(id int) (User, error) {
	var user User
	err := r.db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByFullName(fullName string) (User, error) {
	var user User
	err := r.db.Where("FullName = ?", fullName).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("Email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Store(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
