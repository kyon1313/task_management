package model

type Task struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	TargetDate   string     `json:"target_time"` // DateTime without timestamp ,not sure with this
	PrefRefer    int        `json:"pref_id"`
	Preference   Preference `gorm:"foreignKey:PrefRefer"`
	WorkSpace_ID int        `json:"workspace_id"`
	WorkSpace    Workspace  `gorm:"foreignKey:WorkSpace_ID"`
	Author       int        `json:"user_id"`
	User         User       `gorm:"foreignKey:Author"`
	Is_Deleted   bool       `json:"is_deleted"`
}



	