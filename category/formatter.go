package category

import "time"

type PostFormatter struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

func FormatCreateCategory(category Category) PostFormatter {
	formatter := PostFormatter{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		CreatedAt:         category.CreatedAt,
	}

	return formatter
}

type GetFormatter struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Products          []struct {
		Id        int       `json:"id"`
		Title     string    `json:"title"`
		Price     int       `json:"price"`
		Stock     int       `json:"stock"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"products"`
}

func FormatGetAllCategory(category []Category) []GetFormatter {
	var formatter []GetFormatter

	for _, value := range category {
		formatter = append(formatter, GetFormatter{
			ID:                value.ID,
			Type:              value.Type,
			SoldProductAmount: value.SoldProductAmount,
			CreatedAt:         value.CreatedAt,
			UpdatedAt:         value.UpdatedAt,
		})
	}

	for i, _ := range formatter {
		for _, value := range category[i].Products {
			value := struct {
				Id        int       `json:"id"`
				Title     string    `json:"title"`
				Price     int       `json:"price"`
				Stock     int       `json:"stock"`
				CreatedAt time.Time `json:"created_at"`
				UpdatedAt time.Time `json:"updated_at"`
			}{
				Id:        value.ID,
				Title:     value.Title,
				Price:     value.Price,
				Stock:     value.Stock,
				CreatedAt: value.CreatedAt,
				UpdatedAt: value.UpdatedAt,
			}
			formatter[i].Products = append(formatter[i].Products, value)
		}
	}

	return formatter
}

type PatchFormatter struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func FormatPatchCategory(category Category) PatchFormatter {
	formatter := PatchFormatter{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		UpdatedAt:         category.UpdatedAt,
	}

	return formatter
}

type DeleteFormatter struct {
	Message string `json:"message"`
}

func FormatDeleteCategory() DeleteFormatter {
	formatter := DeleteFormatter{
		Message: "Category has been successfully deleted",
	}

	return formatter
}
