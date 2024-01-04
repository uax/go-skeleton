package utils

import (
	"strconv"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/uax/go-skeleton/config"
)

func JWT() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Config("JWT_SECRET"))},
	})
}

func GenerateToken(id int) fiber.Map {

	jwtTTL, err := strconv.Atoi(config.Config("JWT_TTL"))
	if err != nil {
		return fiber.Map{"error": "JWT_TTL is not a number"}
	}

	claims := jwt.MapClaims{
		"id":  id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * time.Duration(jwtTTL)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return fiber.Map{"error": "Error signing token"}
	}

	return fiber.Map{"token": t}
}
