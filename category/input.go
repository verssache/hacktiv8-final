package category

import "github.com/verssache/hacktiv8-final4/user"

type CreateCategoryInput struct {
	Type string `json:"type" binding:"required"`
	User user.User
}

type FindCategoryInput struct {
	ID int `uri:"id" binding:"required"`
}
