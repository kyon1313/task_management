package route

import (
	"errors"
	"task_management/database"
	"task_management/model"
	"task_management/util"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func CreateUserResponse(user model.User) User {
	return User{
		ID:       user.ID,
		Name:     user.FirstName,
		Lastname: user.Lastname,
	}
}

type Team struct {
	ID          uint   `json:"id"`
	Name        string `json:"team_name"`
	Description string `json:"description"`
	User        User   `json:"owner"`
	Is_Deleted  bool   `json:"is_deleted"`
}

func CreateResponseTeam(team model.Team, user User) Team {
	return Team{
		ID:          team.ID,
		Name:        team.Name,
		Description: team.Description,
		User:        user,
		Is_Deleted:  team.Is_Deleted,
	}
}

func AddTeam(c *fiber.Ctx) error {
	var team model.Team
	new_team := new(model.Team)
	var user model.User

	util.BodyParser(c, &new_team)
	database.DB.Find(&user, "id=?", new_team.Owner)
	database.DB.Find(&team, "name=?", new_team.Name)
	if team.Name == new_team.Name {
		return c.JSON(&fiber.Map{
			"Error":   "Team already exist",
			"success": false,
		})
	}
	if user.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "User not exist!",
		})
	}
	database.DB.Create(&new_team)
	response := CreateResponseTeam(*new_team, CreateUserResponse(user))
	return c.JSON(&fiber.Map{
		"Message": "Team Created",
		"Team":    response,
	})

}

func GetTeams(c *fiber.Ctx) error {
	var teams []model.Team
	database.DB.Find(&teams)
	teamresponse := []Team{}
	if len(teams) == 0 {
		return c.JSON(&fiber.Map{
			"Message": "No Team Exist!",
		})
	}

	for _, team := range teams {
		var user model.User
		database.DB.Find(&user, "id=?", team.Owner)
		response := CreateResponseTeam(team, CreateUserResponse(user))
		teamresponse = append(teamresponse, response)
	}
	return c.JSON(&fiber.Map{
		"Team": teamresponse,
	})
}

func FindTeam(id int, team *model.Team) error {
	database.DB.Find(&team, "id=?", id)
	if team.ID == 0 {
		return errors.New("team not existed")
	}
	return nil
}

func Getteam(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var team model.Team

	database.DB.Find(&team, "id=?", id)
	if team.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Team not exist",
		})

	}
	var user model.User
	database.DB.First(&user, team.Owner)
	response := CreateResponseTeam(team, CreateUserResponse(user))
	return c.JSON(&fiber.Map{
		"Team": response,
	})

}

func DeleteTeam(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var team model.Team

	database.DB.Find(&team, "id=?", id)
	if team.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Team not exist",
		})

	}
	database.DB.Delete(&team)
	return c.JSON(&fiber.Map{
		"Message": "Team Deleted!",
	})

}

func UpdateTeam(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var team model.Team

	database.DB.Find(&team, "id=?", id)
	if team.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Team not exist",
		})

	}
	util.BodyParser(c, &team)
	database.DB.Save(&team)
	var user model.User
	database.DB.First(&user, team.Owner)
	response := CreateResponseTeam(team, CreateUserResponse(user))
	return c.JSON(&fiber.Map{
		"Message": "Updated!",
		"Team":    response,
	})
}
