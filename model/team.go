package model

type Team struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"team_name"`
	Description string `json:"description"`
	Owner       int    `json:"owner"`
	User        User   `gorm:"foreignKey:Owner"`
	Is_Deleted  bool   `json:"is_deleted"`
}
