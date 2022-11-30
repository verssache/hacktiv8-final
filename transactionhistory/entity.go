package transactionhistory

import (
	"github.com/verssache/hacktiv8-final4/product"
	"github.com/verssache/hacktiv8-final4/user"
	"time"
)

type TransactionHistory struct {
	ID         int
	ProductID  int
	UserID     int
	Quantity   int
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    product.Product `gorm:"Constraint:OnDelete:CASCADE;"`
	User       user.User       `gorm:"Constraint:OnDelete:CASCADE;"`
}
