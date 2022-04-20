package model

type Workspace struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	
	UserID     int        `json:"user_id"`
	User       User       `gorm:"foreignKey:UserID"`
	Role       int        `json:"pref_id"`
	Preference Preference `gorm:"foreignKey:Role"`
	TeamID     int        `json:"team_id"`
	Team       Team       `gorm:"foreignKey:TeamID"`
	Is_Deleted bool       `json:"is_deleted"`
}
