package transactionhistory

import "time"

type CreateTransactionHistoryFormatter struct {
	Message         string `json:"message"`
	TransactionBill struct {
		TotalPrice   int    `json:"total_price"`
		Quantity     int    `json:"quantity"`
		ProductTitle string `json:"product_title"`
	} `json:"transaction_bill"`
}

func FormatCreateTransactionHistory(transactionHistory TransactionHistory, productTitle string) CreateTransactionHistoryFormatter {
	formatter := CreateTransactionHistoryFormatter{
		Message: "You have successfully purchased the product",
		TransactionBill: struct {
			TotalPrice   int    `json:"total_price"`
			Quantity     int    `json:"quantity"`
			ProductTitle string `json:"product_title"`
		}{
			TotalPrice:   transactionHistory.TotalPrice,
			Quantity:     transactionHistory.Quantity,
			ProductTitle: productTitle,
		},
	}

	return formatter
}

type MyTransactionFormatter struct {
	Id         int `json:"id"`
	ProductId  int `json:"product_id"`
	UserId     int `json:"user_id"`
	Quantity   int `json:"quantity"`
	TotalPrice int `json:"total_price"`
	Product    struct {
		Id         int       `json:"id"`
		Title      string    `json:"title"`
		Price      int       `json:"price"`
		Stock      int       `json:"stock"`
		CategoryId int       `json:"category_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"product"`
}

func FormatMyTransaction(transactionHistory []TransactionHistory) []MyTransactionFormatter {
	var formatter []MyTransactionFormatter

	for _, value := range transactionHistory {
		formatter = append(formatter, MyTransactionFormatter{
			Id:         value.ID,
			ProductId:  value.ProductID,
			UserId:     value.UserID,
			Quantity:   value.Quantity,
			TotalPrice: value.TotalPrice,
			Product: struct {
				Id         int       `json:"id"`
				Title      string    `json:"title"`
				Price      int       `json:"price"`
				Stock      int       `json:"stock"`
				CategoryId int       `json:"category_id"`
				CreatedAt  time.Time `json:"created_at"`
				UpdatedAt  time.Time `json:"updated_at"`
			}{
				Id:         value.Product.ID,
				Title:      value.Product.Title,
				Price:      value.Product.Price,
				Stock:      value.Product.Stock,
				CategoryId: value.Product.CategoryID,
				CreatedAt:  value.Product.CreatedAt,
				UpdatedAt:  value.Product.UpdatedAt,
			},
		})
	}

	return formatter
}

type AllTransactionHistory struct {
	Id         int `json:"id"`
	ProductId  int `json:"product_id"`
	UserId     int `json:"user_id"`
	Quantity   int `json:"quantity"`
	TotalPrice int `json:"total_price"`
	Product    struct {
		Id         int       `json:"id"`
		Title      string    `json:"title"`
		Price      int       `json:"price"`
		Stock      int       `json:"stock"`
		CategoryId int       `json:"category_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"product"`
	User struct {
		Id        int       `json:"id"`
		Email     string    `json:"email"`
		FullName  string    `json:"full_name"`
		Balance   int       `json:"balance"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"user"`
}

func FormatAllTransactionHistory(transactionHistory []TransactionHistory) []AllTransactionHistory {
	var formatter []AllTransactionHistory

	for _, value := range transactionHistory {
		formatter = append(formatter, AllTransactionHistory{
			Id:         value.ID,
			ProductId:  value.ProductID,
			UserId:     value.UserID,
			Quantity:   value.Quantity,
			TotalPrice: value.TotalPrice,
			Product: struct {
				Id         int       `json:"id"`
				Title      string    `json:"title"`
				Price      int       `json:"price"`
				Stock      int       `json:"stock"`
				CategoryId int       `json:"category_id"`
				CreatedAt  time.Time `json:"created_at"`
				UpdatedAt  time.Time `json:"updated_at"`
			}{
				Id:         value.Product.ID,
				Title:      value.Product.Title,
				Price:      value.Product.Price,
				Stock:      value.Product.Stock,
				CategoryId: value.Product.CategoryID,
				CreatedAt:  value.Product.CreatedAt,
				UpdatedAt:  value.Product.UpdatedAt,
			},
			User: struct {
				Id        int       `json:"id"`
				Email     string    `json:"email"`
				FullName  string    `json:"full_name"`
				Balance   int       `json:"balance"`
				CreatedAt time.Time `json:"created_at"`
				UpdatedAt time.Time `json:"updated_at"`
			}{
				Id:        value.User.ID,
				Email:     value.User.Email,
				FullName:  value.User.FullName,
				Balance:   value.User.Balance,
				CreatedAt: value.User.CreatedAt,
				UpdatedAt: value.User.UpdatedAt,
			},
		})
	}

	return formatter
}
