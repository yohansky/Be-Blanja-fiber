package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/middleware"
	"backend-gin/src/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
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

	user := models.User{
		Name:  data["Name"],
		Email: data["Email"],
		Phone: data["Phone"],
		Store: data["Store"],
		// Password: password,
		RoleId: 1,
	}

	user.SetPassword(data["Password"])

	config.DB.Create(&user)

	return c.JSON(user)
}

func RegisterC(c *fiber.Ctx) error {
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

	user := models.User{
		Name:   data["Name"],
		Email:  data["Email"],
		Phone:  data["Phone"],
		RoleId: 2,
	}

	user.SetPassword(data["Password"])

	config.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	config.DB.Preload("Role").Where("email = ?", data["Email"]).First(&user)

	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "Email not found",
		})
	}

	if err := user.ComparePassword(data["Password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "incorrect Password",
		})
	}

	token, err := middleware.GenerateJwt(strconv.Itoa(int(user.Id)))

	//Cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	roleName := user.Role.Name

	item := map[string]string{
		"Email": data["Email"],
		"Id":    strconv.Itoa(int(user.Id)),
		"Role":  roleName,
		"Token": token,
	}

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(item)
}

type Claims struct {
	jwt.StandardClaims
}

// Otentikasi
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := middleware.ParseJwt(cookie)

	var user models.User

	config.DB.Where("id = ?", id).Preload("Role").First(&user)

	return c.JSON(user)
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
		"Message": "Logout Success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	id, _ := middleware.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	// var user models.User

	user := models.User{
		Id:    uint(userId),
		Name:  data["Name"],
		Email: data["Email"],
		Phone: data["Phone"],
		Store: data["Store"],
	}

	config.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
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

	cookie := c.Cookies("jwt")

	id, _ := middleware.ParseJwt(cookie)

	// var user models.User

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	user.SetPassword(data["Password"])

	config.DB.Model(&user).Updates(user)

	return c.JSON(fiber.Map{
		"Message": "Password has been Changed",
	})
}
