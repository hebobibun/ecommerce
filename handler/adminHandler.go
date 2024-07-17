package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	database "github.com/hebobibun/go-ecommerce/db"
	"github.com/hebobibun/go-ecommerce/helper"
	IF "github.com/hebobibun/go-ecommerce/interfaces"
	model "github.com/hebobibun/go-ecommerce/models"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct{}

var DB = database.NewDB()

func NewAdmin() IF.Admin {
	return &Admin{}
}

func (*Admin) Create(c *fiber.Ctx) error {
	db := DB.OpenDB()
	defer DB.CloseDB(db)

	var admin model.Admin
	if err := c.BodyParser(&admin); err != nil {
		log.Println("Error parsing body: ", err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := helper.ValidateStruct(admin); err != nil {
		log.Println("Error validating struct: ", err)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 14)
	if err != nil {
		log.Println("Error generating hash: ", err)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	admin.Password = string(hash)

	if err := db.Create(&admin).Error; err != nil {
		log.Println("Failed to create admin: ", err)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).SendString("Admin created successfully")
}
