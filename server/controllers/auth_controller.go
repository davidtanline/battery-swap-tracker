package controllers

import (
	"server/database"
	"server/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["Pass"]), 14)

	account := models.Account{
		FirstName:   data["FirstName"],
		LastName:    data["LastName"],
		Email:       data["Email"],
		Pass:        password,
		AccountType: data["AccountType"],
	}

	database.DB.Create(&account)

	return c.JSON(account)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var account models.Account

	database.DB.Where("Email = ?", data["Email"]).First(&account)

	// account not found
	if account.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Account not found.",
		})
	}

	// if account found, check if password is correct
	if err := bcrypt.CompareHashAndPassword(account.Pass, []byte(data["Pass"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password.",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(account.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // expires after 1 hour
	})

	signedString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login.",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    signedString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logged in successfully",
	})
}

func Account(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized.",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var account models.Account

	database.DB.Where("id = ?", claims.Issuer).First(&account)

	return c.JSON(account)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logged out successfully.",
	})
}
