package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	database "github.com/hebobibun/go-ecommerce/db"
	"github.com/hebobibun/go-ecommerce/helper"
	IF "github.com/hebobibun/go-ecommerce/interfaces"
	model "github.com/hebobibun/go-ecommerce/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct{}

func NewUser() IF.User {
	return &User{}
}

func (*User) RegisterUser(c *fiber.Ctx) error {
	db := DB.OpenDB()
	defer DB.CloseDB(db)

	user := model.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := helper.ValidateStruct(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	user.Password = string(hash)

	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).SendString("User created successfully")
}

func (*User) LoginUser(c *fiber.Ctx) error {
	db := DB.OpenDB()
	defer DB.CloseDB(db)

	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	userData := model.User{}
	if err := db.Where("username = ?", user.Username).First(&userData).Error; err != nil {
		return c.Status(404).SendString(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return c.Status(401).SendString(err.Error())
	}

	token, err := helper.GenerateToken("user", user.ID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if err := database.Client.Set(database.Ctx, "user:"+user.ID.String(), token, 0).Err(); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"token": token,
	})
}
