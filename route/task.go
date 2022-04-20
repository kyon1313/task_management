package route

import (
	"task_management/database"
	"task_management/model"
	"task_management/util"

	"github.com/gofiber/fiber/v2"
)

func AddTask(c *fiber.Ctx) error {
	new_task := new(model.Task)
	var task model.Task
	util.BodyParser(c, &new_task)
	database.DB.Find(&task, "title=?", new_task.Title)
	if task.Title == new_task.Title {
		return c.JSON(&fiber.Map{
			"Message": "Task already existed!",
		})
	}
	var pref model.Preference
	database.DB.Find(&pref, "id=?", new_task.PrefRefer)
	if pref.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Preference does not exist!",
		})
	}
	var workspace model.Workspace
	database.DB.Find(&workspace, "id=?", new_task.WorkSpace_ID)
	if workspace.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Workspace does not exist!",
		})
	}
	var user model.User
	database.DB.Find(&user, "id=?", new_task.Author)
	if user.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "User does not exist!",
		})
	}

	Users := CreateUserSerializer(user)
	Prefs := CreatePreferenceSerializer(pref)
	WorkS := CreateWorkspaceSerializer(workspace)
	response := CreateTaskSerializer(*new_task, Prefs, WorkS, Users)

	database.DB.Create(&new_task)
	return c.JSON(&fiber.Map{
		"Task": response,
	})

}

func GetTasks(c *fiber.Ctx) error {
	var task []model.Task
	database.DB.Find(&task)
	responseTask := []TaskSerializer{}
	if len(task) == 0 {
		return c.JSON(&fiber.Map{
			"Message": "There is no task here",
		})
	}

	for _, v := range task {
		var pref model.Preference
		var workspace model.Workspace
		var user model.User

		database.DB.Find(&pref, "id=?", v.PrefRefer)
		database.DB.Find(&workspace, "id=?", v.WorkSpace_ID)
		database.DB.Find(&user, "id=?", v.Author)

		Users := CreateUserSerializer(user)
		Prefs := CreatePreferenceSerializer(pref)
		WorkS := CreateWorkspaceSerializer(workspace)
		response := CreateTaskSerializer(v, Prefs, WorkS, Users)

		responseTask = append(responseTask, response)

	}

	return c.JSON(&fiber.Map{
		"Task": responseTask,
	})
}

func Gettask(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var task model.Task
	database.DB.Find(&task, "id=?", id)
	if task.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Task not exist!",
		})

	}
	var pref model.Preference
	var user model.User
	var workspace model.Workspace

	database.DB.First(&pref, task.PrefRefer)
	database.DB.First(&user, task.Author)
	database.DB.First(&workspace, task.WorkSpace_ID)

	Users := CreateUserSerializer(user)
	Prefs := CreatePreferenceSerializer(pref)
	WorkS := CreateWorkspaceSerializer(workspace)
	response := CreateTaskSerializer(task, Prefs, WorkS, Users)

	return c.JSON(&fiber.Map{
		"TAsk": response,
	})

}

func DeleteTask(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var task model.Task
	database.DB.Find(&task, "id=?", id)
	if task.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Task not exist!",
		})

	}
	database.DB.Delete(&task)
	return c.JSON(&fiber.Map{
		"Message": "Task Deleted",
	})
}

func UpdateTask(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var task model.Task
	database.DB.Find(&task, "id=?", id)
	if task.ID == 0 {
		return c.JSON(&fiber.Map{
			"Message": "Task not exist!",
		})

	}
	util.BodyParser(c, &task)
	database.DB.Save(&task)

	var pref model.Preference
	var user model.User
	var workspace model.Workspace

	database.DB.First(&pref, task.PrefRefer)
	database.DB.First(&user, task.Author)
	database.DB.First(&workspace, task.WorkSpace_ID)

	Users := CreateUserSerializer(user)
	Prefs := CreatePreferenceSerializer(pref)
	WorkS := CreateWorkspaceSerializer(workspace)
	response := CreateTaskSerializer(task, Prefs, WorkS, Users)

	return c.JSON(&fiber.Map{
		"TAsk":         response,
		"TAsk Updated": true,
	})

}
