package product

import "time"

type CreateProductFormatter struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatCreateProduct(product Product) CreateProductFormatter {
	formatter := CreateProductFormatter{
		ID:         product.ID,
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt,
	}

	return formatter
}

type GetProductFormatter struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatGetProduct(product Product) GetProductFormatter {
	formatter := GetProductFormatter{
		Id:         product.ID,
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt,
	}

	return formatter
}

func FormatGetAllProduct(product []Product) []GetProductFormatter {
	var formatter []GetProductFormatter

	for _, value := range product {
		formatter = append(formatter, GetProductFormatter{
			Id:         value.ID,
			Title:      value.Title,
			Price:      value.Price,
			Stock:      value.Stock,
			CategoryID: value.CategoryID,
			CreatedAt:  value.CreatedAt,
		})
	}

	return formatter
}

type UpdateProductFormatter struct {
	Products struct {
		Id         int       `json:"id"`
		Title      string    `json:"title"`
		Price      int       `json:"price"`
		Stock      int       `json:"stock"`
		CategoryId int       `json:"category_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"products"`
}

func FormatUpdateProduct(product Product) UpdateProductFormatter {
	formatter := UpdateProductFormatter{
		Products: struct {
			Id         int       `json:"id"`
			Title      string    `json:"title"`
			Price      int       `json:"price"`
			Stock      int       `json:"stock"`
			CategoryId int       `json:"category_id"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
		}{
			Id:         product.ID,
			Title:      product.Title,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryId: product.CategoryID,
			CreatedAt:  product.CreatedAt,
			UpdatedAt:  product.UpdatedAt,
		},
	}

	return formatter
}

type DeleteProductFormatter struct {
	Message string `json:"message"`
}

func FormatDeleteProduct() DeleteProductFormatter {
	formatter := DeleteProductFormatter{
		Message: "Product has been successfully deleted",
	}

	return formatter
}
