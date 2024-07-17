package middleware

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	database "github.com/hebobibun/go-ecommerce/db"
)

func JWT(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" {
		return c.Status(401).SendString("Unauthorized - No Authorization header")
	}

	token := strings.Split(auth, " ")
	if len(token) != 2 {
		return c.Status(401).SendString("Unauthorized - Invalid Authorization header")
	}

	jwtToken, err := jwt.Parse(token[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return c.Status(401).SendString("Unauthorized - Invalid token")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).SendString("Unauthorized - Invalid token claims")
	}

	if _, ok := claims["role"]; !ok {
		return c.Status(401).SendString("Unauthorized - Invalid token role")
	} else if claims["role"] != "admin" {
		return c.Status(401).SendString("Unauthorized - Invalid token role - " + claims["role"].(string))
	}

	ctx := context.Background()
	isBlacklisted, err := database.Client.Get(ctx, "admin:"+claims["sub"].(string)).Result()
	if err != nil {
		return c.Status(401).SendString("Unauthorized - Invalid token sub")
	}

	if isBlacklisted == "true" {
		return c.Status(401).SendString("Unauthorized - Invalid token sub - " + claims["sub"].(string))
	}

	c.Locals("adminID", claims["sub"])

	return c.Next()
}
