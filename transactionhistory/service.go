package transactionhistory

type Service interface {
	Store(input CreateTransactionInput, UserID, TotalPrice int) (TransactionHistory, error)
	FindAll() ([]TransactionHistory, error)
	FindByID(id int) (TransactionHistory, error)
	FindAllByID(id int) ([]TransactionHistory, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input CreateTransactionInput, UserID, TotalPrice int) (TransactionHistory, error) {
	transactionHistory := TransactionHistory{}
	transactionHistory.UserID = UserID
	transactionHistory.ProductID = input.ProductID
	transactionHistory.Quantity = input.Quantity
	transactionHistory.TotalPrice = TotalPrice

	transactionHistory, err := s.repository.Store(transactionHistory)
	if err != nil {
		return transactionHistory, err
	}

	return transactionHistory, nil
}

func (s *service) FindAll() ([]TransactionHistory, error) {
	transactionHistories, err := s.repository.FindAll()
	if err != nil {
		return transactionHistories, err
	}

	return transactionHistories, nil
}

func (s *service) FindByID(id int) (TransactionHistory, error) {
	transactionHistory, err := s.repository.FindByID(id)
	if err != nil {
		return transactionHistory, err
	}

	return transactionHistory, nil
}

func (s *service) FindAllByID(id int) ([]TransactionHistory, error) {
	transactionHistories, err := s.repository.FindAllByID(id)
	if err != nil {
		return transactionHistories, err
	}

	return transactionHistories, nil
}
