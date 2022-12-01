package transactionhistory

import "gorm.io/gorm"

type Repository interface {
	Store(transactionHistory TransactionHistory) (TransactionHistory, error)
	FindAll() ([]TransactionHistory, error)
	FindByID(id int) (TransactionHistory, error)
	FindAllByID(id int) ([]TransactionHistory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(transactionHistory TransactionHistory) (TransactionHistory, error) {
	err := r.db.Create(&transactionHistory).Error
	if err != nil {
		return transactionHistory, err
	}

	return transactionHistory, nil
}

func (r *repository) FindAll() ([]TransactionHistory, error) {
	var transactionHistories []TransactionHistory
	err := r.db.Preload("User").Preload("Product").Find(&transactionHistories).Error
	if err != nil {
		return transactionHistories, err
	}

	return transactionHistories, nil
}

func (r *repository) FindByID(id int) (TransactionHistory, error) {
	var transactionHistory TransactionHistory
	err := r.db.Preload("User").Preload("Product").Where("ID = ?", id).First(&transactionHistory).Error
	if err != nil {
		return transactionHistory, err
	}

	return transactionHistory, nil
}

func (r *repository) FindAllByID(id int) ([]TransactionHistory, error) {
	var transactionHistories []TransactionHistory
	err := r.db.Preload("User").Preload("Product").Where("user_id = ?", id).Find(&transactionHistories).Error
	if err != nil {
		return transactionHistories, err
	}

	return transactionHistories, nil
}
