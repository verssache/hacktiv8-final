package user

import (
	"strconv"
	"time"
)

type RegisterUserFormatter struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatRegisterUser(user User) RegisterUserFormatter {
	formatter := RegisterUserFormatter{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}

	return formatter
}

type LoginUserFormatter struct {
	Token string `json:"token"`
}

func FormatLoginUser(token string) LoginUserFormatter {
	formatter := LoginUserFormatter{
		Token: token,
	}

	return formatter
}

type TopUpBalanceFormatter struct {
	Message string `json:"message"`
}

func FormatTopUpBalance(totalBalance int) TopUpBalanceFormatter {
	formatter := TopUpBalanceFormatter{
		Message: "Your balance has been successfully updated to Rp " + strconv.Itoa(totalBalance),
	}

	return formatter
}
