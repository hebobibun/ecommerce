package database

import (
	"log"

	model "github.com/hebobibun/go-ecommerce/models"
)

var DB = NewDB()

func Migrate() {
	db := DB.OpenDB()
	defer DB.CloseDB(db)

	err := DB.OpenDB().AutoMigrate(&model.User{})
	errHandler(err)
	err = DB.OpenDB().AutoMigrate(&model.Admin{})
	errHandler(err)
	err = DB.OpenDB().AutoMigrate(&model.Products{})
	errHandler(err)
	err = DB.OpenDB().AutoMigrate(&model.Categories{})
	errHandler(err)
	err = DB.OpenDB().AutoMigrate(&model.CartItems{})
	errHandler(err)
	err = DB.OpenDB().AutoMigrate(&model.Orders{})
	errHandler(err)
	err = DB.OpenDB().AutoMigrate(&model.OrderItems{})
	errHandler(err)
}

func errHandler(err error) {
	if err != nil {
		log.Print(err)
	}
}
