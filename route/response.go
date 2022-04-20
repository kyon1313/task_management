package route

import "task_management/model"

type UserSerializer struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func CreateUserSerializer(user model.User) UserSerializer {
	return UserSerializer{
		ID:        user.ID,
		Firstname: user.FirstName,
		Lastname:  user.Lastname,
	}
}

type PreferenceSerializer struct {
	ID       uint   `json:"id"`
	Ref_type string `json:"ref_type"`
	Title    string `json:"title"`
}

func CreatePreferenceSerializer(pref model.Preference) PreferenceSerializer {
	return PreferenceSerializer{
		ID:       pref.ID,
		Ref_type: pref.Ref_type,
		Title:    pref.Title,
	}
}

type TeamSerializer struct {
	ID          uint           `json:"id"`
	Name        string         `json:"team_name"`
	Description string         `json:"description"`
	User        UserSerializer `json:"owner"`
	Is_Deleted  bool           `json:"is_deleted"`
}

func CreateTeamSerializer(team model.Team, user UserSerializer) TeamSerializer {
	return TeamSerializer{
		ID:          team.ID,
		Name:        team.Name,
		Description: team.Description,
		User:        user,
		Is_Deleted:  team.Is_Deleted,
	}
}

type WorkSpaceSerializer struct {
	ID uint `json:"id"`
}

func CreateWorkspaceSerializer(workspace model.Workspace) WorkSpaceSerializer {

	return WorkSpaceSerializer{
		ID: workspace.ID,
	}
}

type TaskSerializer struct {
	ID          uint                 `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	TargerDate  string               `json:"target_date"` // time without timestamp ,dunno what to do e so i put string muna
	Severity    PreferenceSerializer `json:"Severity"`
	WorkspaceId WorkSpaceSerializer  `json:"workspace_id"`
	Author      UserSerializer       `json:"author"`
	Is_Deleted  bool                 `json:"is_deleted"`
}

func CreateTaskSerializer(task model.Task, pref PreferenceSerializer, workspace WorkSpaceSerializer, user UserSerializer) TaskSerializer {
	return TaskSerializer{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		TargerDate:  task.TargetDate,
		Severity:    pref,
		WorkspaceId: workspace,
		Author:      user,
		Is_Deleted:  task.Is_Deleted,
	}
}
