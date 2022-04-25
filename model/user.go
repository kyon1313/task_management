package model

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
	
	Username  string `json:"username"`
	Password  string `json:"password"`
}
