package model

type Preference struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Ref_type string `json:"ref_type"`
	Title    string `json:"title"`
}
