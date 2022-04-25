package route

import (
	"fmt"

	"task_management/database"
	"task_management/model"
	"task_management/util"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(c *fiber.Ctx) error {
	var user model.User
	new_user := new(model.User)
	util.BodyParser(c, new_user)
	//email does not exist

	database.DB.Find(&user, "username=?", new_user.Username)

	if user.Username == new_user.Username {
		return c.JSON(&fiber.Map{
			"message": "Username already exist",
			"error":   "Registration Failed",
		})
	} else {
		hash, _ := HashPassword(new_user.Password)
		new_user.Password = hash

		database.DB.Create(new_user)
		fmt.Println(new_user)
	}

	return c.JSON(&fiber.Map{
		"Success": true,
		"User":    new_user,
	})

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func GetUsers(c *fiber.Ctx) error {
	var user []model.User
	database.DB.Find(&user)

	if len(user) == 0 {
		return c.JSON(&fiber.Map{
			"Message": "No Data",
		})
	}
	return c.JSON(&fiber.Map{
		"User": user,
	})
}

func GetUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var user model.User
	database.DB.Find(&user, "id=?", id)
	if user.ID == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"Message": "User not exist!",
		})

	}

	return c.JSON(&fiber.Map{
		"User": user,
	})

}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var user model.User
	database.DB.Find(&user, "id=?", id)
	if user.ID == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"Message": "User not exist!",
		})
	}
	database.DB.Delete(user)
	return c.JSON(&fiber.Map{
		"MEssage": "User Deleted",
	})

}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var user model.User
	database.DB.Find(&user, "id=?", id)
	if user.ID == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"Message": "User not exist!",
		})
	}
	util.BodyParser(c, &user)

	hash, _ := HashPassword(user.Password)
	user.Password = hash
	database.DB.Save(user)

	return c.JSON(&fiber.Map{
		"User": user,
	})
}
