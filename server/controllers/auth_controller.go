package controllers

import (
	"server/database"
	"server/models"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Next(err)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["pass"]), 14)

	account := models.Account{
		FirstName:   data["FirstName"],
		LastName:    data["LastName"],
		Email:       data["Email"],
		Pass:        password,
		AccountType: data["AccountType"],
	}

	database.DB.Create(&account)

	c.JSON(account)
}
