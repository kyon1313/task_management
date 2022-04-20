package route

import (
	"task_management/database"
	"task_management/model"
	"task_management/util"

	"github.com/gofiber/fiber/v2"
)

func AddPref(c *fiber.Ctx) error {
	var pref model.Preference

	util.BodyParser(c, &pref)
	database.DB.Create(&pref)
	return c.JSON(&fiber.Map{
		"Pref": pref,
	})
	
}

func GetPrefs(c *fiber.Ctx) error {
	var pref []model.Preference
	database.DB.Find(&pref)
	if len(pref) == 0 {
		return c.JSON(&fiber.Map{
			"Message": "No Preference!",
		})
	}
	return c.JSON(&fiber.Map{
		"Preferences": pref,
	})
}

func Getpref(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var pref model.Preference
	database.DB.Find(&pref, "id=?", id)
	if pref.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "No Preference!",
		})
	}
	return c.JSON(&fiber.Map{
		"Preferences": pref,
	})
}

func DeletePref(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var pref model.Preference
	database.DB.Find(&pref, "id=?", id)
	if pref.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "No Preference!",
		})
	}
	database.DB.Delete(&pref)
	return c.JSON(&fiber.Map{
		"Message": "Pref successfully deleted!",
		"Delete":  true,
	})
}

func UpdatePref(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var pref model.Preference
	database.DB.Find(&pref, "id=?", id)
	if pref.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "No Preference!",
		})
	}
	util.BodyParser(c, &pref)
	database.DB.Save(&pref)
	return c.JSON(&fiber.Map{
		"Message":    "Pref successfully Updated!",
		"Preference": pref,
	})
}
