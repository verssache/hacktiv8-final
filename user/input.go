package user

type RegisterUserInput struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email" gorm:"unique_index"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type TopUpBalanceInput struct {
	Balance int `json:"balance" binding:"required"`
}
