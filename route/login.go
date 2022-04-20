package route

import (
	"fmt"
	"log"
	"task_management/database"
	"task_management/model"
	"task_management/util"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Log(c *fiber.Ctx) error {
	var log Login
	var user model.User
	util.BodyParser(c, &log)
	database.DB.Find(&user, "username=?", log.Username)
	if log.Username != user.Username {
		return c.JSON(&fiber.Map{
			"Message":       "Wrong Username or Password",
			"Login success": false,
		})
	} else {
		match := CheckPasswordHash([]byte(user.Password), []byte(log.Password))
		if !match {
			return c.JSON(&fiber.Map{
				"Message":       "Wrong Username or Password",
				"Login success": false,
			})
		}
		fmt.Print("password match", match)
		return c.JSON(&fiber.Map{
			"Welcome":       user.FirstName,
			"Login success": true,
		})
	}

}

func CheckPasswordHash(hash []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		log.Println("Unable to compare password", err)
		return false
	}
	return true
}
