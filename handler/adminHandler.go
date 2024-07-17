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

func (*Admin) Login(c *fiber.Ctx) error {
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

	user := model.Admin{}
	if err := db.Where("username = ?", admin.Username).First(&user).Error; err != nil {
		log.Println("Error finding user: ", err)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(admin.Password)); err != nil {
		log.Println("Error comparing passwords: ", err)
		return c.Status(http.StatusUnauthorized).SendString(err.Error())
	}

	token, err := helper.GenerateToken("admin", user.ID)
	if err != nil {
		log.Println("Error generating token: ", err)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if err := database.Client.Set(database.Ctx, "admin:"+user.ID.String(), token, 0).Err(); err != nil {
		log.Println("Error storing token in Redis: ", err)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	log.Println("stored to redis:", database.Client.Get(database.Ctx, "admin:"+user.ID.String()))

	return c.Status(http.StatusOK).JSON(map[string]string{
		"token": token,
	})
}
