package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hebobibun/go-ecommerce/helper"
	IF "github.com/hebobibun/go-ecommerce/interfaces"
	model "github.com/hebobibun/go-ecommerce/models"
)

type Product struct{}

func NewProduct() IF.Product {
	return &Product{}
}

func (*Product) CreateProduct(c *fiber.Ctx) error {
	db := DB.OpenDB()
	defer DB.CloseDB(db)

	var product model.Products
	if err := c.BodyParser(&product); err != nil {
		log.Println("Error parsing body: ", err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err := helper.ValidateStruct(product)
	if err != nil {
		log.Println("Error validating struct: ", err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := db.Create(&product).Error; err != nil {
		log.Println("Failed to create product: ", err)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).SendString("Product created successfully")
}

func (*Product) CreateCategory(c *fiber.Ctx) error {
	db := DB.OpenDB()
	defer DB.CloseDB(db)

	var category model.Categories
	if err := c.BodyParser(&category); err != nil {
		log.Println("Error parsing body: ", err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err := helper.ValidateStruct(category)
	if err != nil {
		log.Println("Error validating struct: ", err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := db.Create(&category).Error; err != nil {
		log.Println("Failed to create category: ", err)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).SendString("Category created successfully")
}
