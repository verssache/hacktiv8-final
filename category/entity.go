package category

import (
	"github.com/verssache/hacktiv8-final4/product"
	"time"
)

type Category struct {
	ID                int
	Type              string
	SoldProductAmount int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Products          []product.Product `gorm:"Constraint:OnDelete:CASCADE;"`
}
