package product

import "github.com/verssache/hacktiv8-final4/user"

type CreateProductInput struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
	User       user.User
}

type FindProductInput struct {
	ID int `uri:"id" binding:"required"`
}
