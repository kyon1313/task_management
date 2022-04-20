package route

import (
	"task_management/database"
	"task_management/model"
	"task_management/util"

	"github.com/gofiber/fiber/v2"
)

type Workspace struct {
	ID         uint             `json:"id"`
	User       User             `json:"user"`
	Pref       model.Preference `json:"preference"`
	Team       Team             `json:"team"`
	Is_Deleted bool             `json:"is_deleted"`
}

func CreateWorkSpace(workspace model.Workspace, user User, team Team, pref model.Preference) Workspace {

	return Workspace{
		ID:         workspace.ID,
		User:       user,
		Pref:       pref,
		Team:       team,
		Is_Deleted: workspace.Is_Deleted,
	}
}

func AddWorkspace(c *fiber.Ctx) error {
	var workspace model.Workspace
	util.BodyParser(c, &workspace)
	//check user is it exist
	var user model.User
	database.DB.Find(&user, "id=?", workspace.UserID)
	if user.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "User not Exist",
		})
	}
	var pref model.Preference
	database.DB.Find(&pref, "id=?", workspace.Role)
	if pref.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Preference not Exist",
		})
	}
	var team model.Team
	database.DB.Find(&team, "id=?", workspace.TeamID)
	if team.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Team not Exist",
		})
	}
	database.DB.Create(&workspace)

	response := CreateWorkSpace(workspace, CreateUserResponse(user), CreateResponseTeam(team, CreateUserResponse(user)), pref)
	return c.JSON(&fiber.Map{
		"Success":   true,
		"Workspace": response,
	})
}

func GetWorkSpaces(c *fiber.Ctx) error {
	var workspace []model.Workspace
	database.DB.Find(&workspace)
	responseWorkspace := []Workspace{}
	if len(workspace) == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Workspace is empty!",
		})
	}
	for _, v := range workspace {
		var user model.User
		var pref model.Preference
		var team model.Team

		database.DB.Find(&user, "id=?", v.UserID)
		database.DB.Find(&pref, "id=?", v.Role)
		database.DB.Find(&team, "id=?", v.TeamID)
		response := CreateWorkSpace(v, CreateUserResponse(user), CreateResponseTeam(team, CreateUserResponse(user)), pref)
		responseWorkspace = append(responseWorkspace, response)
	}

	return c.JSON(&fiber.Map{
		"WorkSpaces": responseWorkspace,
	})

}

func GetWorkspace(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var workspace model.Workspace

	database.DB.Find(&workspace, "id=?", id)
	if workspace.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Workspace not exist",
		})

	}

	var user model.User
	var team model.Team
	var pref model.Preference

	database.DB.First(&user, workspace.UserID)
	database.DB.First(&pref, workspace.Role)
	database.DB.First(&team, workspace.TeamID)

	response := CreateWorkSpace(workspace, CreateUserResponse(user), CreateResponseTeam(team, CreateUserResponse(user)), pref)
	return c.JSON(&fiber.Map{
		"Workspace": response,
	})
}

func DeleteWorkspace(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var workspace model.Workspace

	database.DB.Find(&workspace, "id=?", id)
	if workspace.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Workspace not exist",
		})

	}
	database.DB.Delete(&workspace)
	return c.JSON(&fiber.Map{
		"Message": "Workspace Deleted",
	})

}

func UpdateWorkSpace(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var workspace model.Workspace

	database.DB.Find(&workspace, "id=?", id)
	if workspace.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Workspace not exist",
		})

	}
	util.BodyParser(c, &workspace)
	database.DB.Save(&workspace)

	var user model.User
	var team model.Team
	var pref model.Preference

	database.DB.First(&user, workspace.UserID)
	database.DB.First(&pref, workspace.Role)
	database.DB.First(&team, workspace.TeamID)

	response := CreateWorkSpace(workspace, CreateUserResponse(user), CreateResponseTeam(team, CreateUserResponse(user)), pref)
	return c.JSON(&fiber.Map{
		"Workspace": response,
	})
}
