package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//validasi jika password tidak amtch confirm password
	if data["Password"] != data["Passwordconfirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "Password do not match",
		})
	}

	//hashed pass
	password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	user := models.User{
		Name:     data["Name"],
		Email:    data["Email"],
		Phone:    data["Phone"],
		Store:    data["Store"],
		Password: password,
		Role:     "Seller",
	}

	config.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	config.DB.Where("email = ?", data["Email"]).First(&user)

	if user.ID == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "Email not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["Password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "incorrect Password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte("Secret"))

	item := map[string]string{
		"Email": data["Email"],
		"Role":  user.Role,
		"Token": token,
	}

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(item)
}
