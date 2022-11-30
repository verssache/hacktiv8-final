package helper

import (
	"github.com/verssache/hacktiv8-final4/category"
	"github.com/verssache/hacktiv8-final4/product"
	"github.com/verssache/hacktiv8-final4/transactionhistory"
	"github.com/verssache/hacktiv8-final4/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	cfg := LoadConfig()
	dsn := cfg.Database.User + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ":" + cfg.Database.Port + ")/" + cfg.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&user.User{}, &product.Product{}, &category.Category{}, &transactionhistory.TransactionHistory{})
	return db
}
